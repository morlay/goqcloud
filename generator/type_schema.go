package generator

import (
	"sort"
	"strings"

	"github.com/go-courier/codegen"
)

type TypeSchema struct {
	ImportPath     string
	Name           string
	Desc           string
	Type           *BasicType
	Items          *TypeSchema
	AdditionalProp *TypeSchema
	AllOf          []*TypeSchema
	Props          map[string]*TypeSchema
	Required       bool
	Tag            string
}

func (s *TypeSchema) IsRef() bool {
	return s.Type == nil && !s.IsArray() && !s.IsMap() && !s.IsObject()
}

func (s *TypeSchema) IsArray() bool {
	return s.Items != nil
}

func (s *TypeSchema) IsObject() bool {
	return s.Props != nil || len(s.AllOf) > 0
}

func (s *TypeSchema) IsMap() bool {
	return s.AdditionalProp != nil
}

func (s *TypeSchema) Write(file *codegen.File) {
	file.WriteBlock(
		codegen.Comments(s.Desc),
		codegen.DeclType(
			codegen.Var(s.SnippetType(file), s.Name),
		),
	)
}

func (s *TypeSchema) SnippetType(file *codegen.File) codegen.SnippetType {
	if s.IsArray() {
		return codegen.Slice(s.Items.SnippetType(file))
	}

	if s.IsObject() {
		fields := make([]*codegen.SnippetField, 0)

		for _, embedType := range s.AllOf {
			fields = append(fields, codegen.Var(embedType.SnippetType(file)))
		}

		names := make([]string, 0)
		for name := range s.Props {
			names = append(names, name)
		}
		sort.Strings(names)

		for _, name := range names {
			prop := s.Props[name]

			tags := map[string][]string{}
			tag := s.Tag
			if tag == "" {
				tag = "json"
			}
			tags[tag] = append(tags[tag], name)
			if !prop.Required {
				tags[tag] = append(tags[tag], "omitempty")
			}

			fields = append(fields, codegen.Var(prop.SnippetType(file), name).WithTags(tags).WithComments(prop.Desc))
		}

		return codegen.Struct(fields...)
	}

	maybeStarType := func(tpe codegen.SnippetType) codegen.SnippetType {
		if !s.Required {
			return codegen.Star(tpe)
		}
		return tpe
	}

	if s.Name != "" {
		if s.ImportPath != "" {
			return maybeStarType(codegen.Type(file.Use(s.ImportPath, s.Name)))
		}
		return maybeStarType(codegen.Type(s.Name))

	}

	if s.Type != nil {
		switch *s.Type {
		case BasicTypeBoolean:
			return maybeStarType(codegen.Bool)
		case BasicTypeString:
			return maybeStarType(codegen.String)
		case BasicTypeInteger:
			return maybeStarType(codegen.Int64)
		case BasicTypeFloat:
			return maybeStarType(codegen.Float64)
		case BasicTypeTimestamp, BasicTypeDate:
			return maybeStarType(codegen.Type(file.Use("time", "Time")))
		}
	}

	return codegen.Interface()
}

func (s *TypeSchema) AddProp(name string, propSchema *TypeSchema) {
	if s.Props == nil {
		s.Props = map[string]*TypeSchema{}
	}
	if s.Props[name] == nil {
		s.Props[name] = propSchema
	}
}

type BasicType string

const (
	BasicTypeDate      BasicType = "Date"
	BasicTypeTimestamp BasicType = "Timestamp"
	BasicTypeString    BasicType = "String"
	BasicTypeInteger   BasicType = "Integer"
	BasicTypeBoolean   BasicType = "Boolean"
	BasicTypeFloat     BasicType = "Float"
)

const arrayPrefix = "Array of "

func IsArrayType(tpe string) bool {
	return strings.HasPrefix(tpe, arrayPrefix)
}

func IndirectType(tpe string) (basicType *BasicType) {
	t := BasicType(strings.TrimLeft(tpe, arrayPrefix))
	return &t
}

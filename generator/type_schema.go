package generator

import (
	"sort"
	"strings"

	"github.com/morlay/goqcloud/generator/codegen"
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
	file.WriteComments(s.Desc)
	file.WriteTypeSpec(s.Name, s.WriteType)
}

func (s *TypeSchema) WriteType(file *codegen.File) {
	if s.IsArray() {
		file.WriteString("[]")
		s.Items.WriteType(file)
		return
	}

	if s.IsObject() {
		file.WriteStructType(func(file *codegen.File) {
			for _, embedType := range s.AllOf {
				embedType.WriteType(file)
				file.WriteLine()
			}

			names := make([]string, 0)
			for name := range s.Props {
				names = append(names, name)
			}
			sort.Strings(names)

			for _, name := range names {
				prop := s.Props[name]

				file.WriteComments(prop.Desc)

				file.WriteField(name, prop.WriteType)

				tags := map[string][]string{}

				tag := s.Tag
				if tag == "" {
					tag = "json"
				}

				tags[tag] = append(tags[tag], name)

				if !prop.Required {
					tags[tag] = append(tags[tag], "omitempty")
				}

				file.WriteTags(tags)
				file.WriteLine()
			}
		})
		return
	}

	ptrMark := "*"
	if s.Required {
		ptrMark = ""
	}

	if s.Name != "" {
		if s.ImportPath != "" {
			file.WriteString(file.Use(s.ImportPath, s.Name))
			return
		}
		file.WriteString(ptrMark + s.Name)
		return
	}

	if s.Type != nil {
		switch *s.Type {
		case BasicTypeBoolean:
			file.WriteString(ptrMark + "bool")
		case BasicTypeString:
			file.WriteString(ptrMark + "string")
		case BasicTypeInteger:
			file.WriteString(ptrMark + "int64")
		case BasicTypeFloat:
			file.WriteString(ptrMark + "float64")
		case BasicTypeTimestamp:
			file.WriteString(ptrMark + file.Use("time", "Time"))
		}
	}

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

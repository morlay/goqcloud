package codegen

import (
	"bytes"
	"fmt"
	"go/build"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func NewFile(pkgName string) *File {
	return &File{
		PkgName: Identifier(pkgName).LowerCamelCase(),
	}
}

type File struct {
	PkgName  string
	Filename string
	imports  map[string]string
	bytes.Buffer
}

func (file *File) SetFilename(filename string) {
	file.Filename = filename
}

func (file *File) WriteFile() (int, error) {
	dir := filepath.Dir(file.Filename)

	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return -1, err
		}
	}

	f, err := os.Create(file.Filename)
	defer f.Close()
	if err != nil {
		return -1, err
	}

	n3, err := f.Write(file.Bytes())
	if err != nil {
		return -1, err
	}

	if err := f.Sync(); err != nil {
		return -1, err
	}

	return n3, nil
}

func (file *File) Bytes() []byte {
	buf := &bytes.Buffer{}

	buf.WriteString(`package ` + Identifier(file.PkgName).LowerLinkCase() + `
`)

	if file.imports != nil {
		buf.WriteString(`import (
`)
		for importPath, alias := range file.imports {
			buf.WriteString(alias)
			buf.WriteString(" ")
			buf.WriteString(strconv.Quote(importPath))
			buf.WriteString("\n")
		}

		buf.WriteString(`)
`)
	}

	io.Copy(buf, &file.Buffer)

	return Format(file.Filename, buf.Bytes())
}

func Format(filename string, src []byte) []byte {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, filename, src, parser.ParseComments)
	if err != nil {
		for i, line := range bytes.Split(src, []byte("\n")) {
			fmt.Printf("\t%d\t%s\n", i+1, line)
		}
		panic(fmt.Errorf("go codes parse failed: %s in %s", err.Error(), filename))
	}
	buf := &bytes.Buffer{}
	if err := format.Node(buf, fileSet, file); err != nil {
		for i, line := range bytes.Split(src, []byte("\n")) {
			fmt.Printf("\t%d\t%s\n", i+1, line)
		}
		panic(fmt.Errorf("go codes format failed: %s in %s", err.Error(), filename))
	}
	return buf.Bytes()
}

func (file *File) WriteComments(desc string) {
	list := strings.Split(desc, "\n")
	for i := range list {
		if list[i] != "" {
			file.WriteString("// " + list[i] + "\n")
		}
	}
}

func (file *File) TypeWriter(tpe string) func(file *File) {
	return func(file *File) {
		file.WriteString(tpe)
	}
}

func (file *File) WriteCall(name string, args ...string) {
	file.WriteString(".")
	file.WriteString(name)
	file.WriteString("(")

	for i, arg := range args {
		if i > 0 {
			file.WriteString(", ")
		}
		file.WriteString(arg)
	}

	file.WriteString(")")
}

func (file *File) WriteReturn(args ...string) {
	file.WriteString("return ")

	for i, arg := range args {
		if i > 0 {
			file.WriteString(", ")
		}
		file.WriteString(arg)
	}

	file.WriteLine()
}

func (file *File) WriteTypeSpec(name string, writeType func(file *File)) {
	file.WriteString("type")
	file.WriteSpace()

	file.WriteField(name, writeType)
}

func (file *File) WriteFuncSpec(
	name string,
	writeReceiver func(file *File),
	writeParams func(file *File),
	writeReturn func(file *File),
	writeBlock func(file *File),
) {
	file.WriteString("func")
	file.WriteSpace()

	if writeReceiver != nil {
		file.WriteString("(")
		writeReceiver(file)
		file.WriteString(")")
		file.WriteSpace()
	}

	file.WriteString(name)
	file.WriteString("(")
	if writeParams != nil {
		writeParams(file)
	}
	file.WriteString(")")

	if writeReturn != nil {
		file.WriteString("(")
		writeReturn(file)
		file.WriteString(")")
	}

	if writeBlock != nil {
		file.WriteString("{")
		file.WriteLine()

		writeBlock(file)

		file.WriteString("}")
		file.WriteLine()
	}
}

func (file *File) WriteField(name string, writeType func(file *File)) {
	file.WriteString(name)
	file.WriteSpace()
	writeType(file)
}

func (file *File) WriteStructType(writeFields func(file *File)) {
	file.WriteString("struct {")
	file.WriteLine()

	writeFields(file)

	file.WriteString("}")
	file.WriteLine()
}

func (file *File) WriteSpace() {
	file.WriteString(" ")
}

func (file *File) WriteComma() {
	file.WriteString(", ")
}

func (file *File) WriteLine() {
	file.WriteString("\n")
}

func (file *File) WriteTags(tags map[string][]string) {

	file.WriteString("`")

	count := 0
	for k, valueAndFlags := range tags {
		if count > 0 {
			file.WriteString(" ")
		}

		file.WriteString(k)
		file.WriteString(":")
		file.WriteString(strconv.Quote(strings.Join(valueAndFlags, ",")))

		count++
	}

	file.WriteString("`")
}

func (file *File) Use(importPath string, exposedName string) string {
	importPath = deVendor(importPath)

	if file.imports == nil {
		file.imports = map[string]string{}
	}

	if file.imports[importPath] == "" {
		pkg, err := build.Import(importPath, "", build.ImportComment)
		if err != nil {
			panic(err)
		}
		file.imports[pkg.ImportPath] = Identifier(importPath).LowerSnakeCase()
	}

	return file.imports[importPath] + "." + exposedName
}

func (file *File) Type(tpe reflect.Type) string {
	if tpe.PkgPath() != "" {
		return file.Use(tpe.PkgPath(), tpe.Name())
	}

	switch tpe.Kind() {
	case reflect.Array:
		return fmt.Sprintf("[%d]%s", tpe.Len(), file.Type(tpe.Elem()))
	case reflect.Slice:
		return fmt.Sprintf("[]%s", file.Type(tpe.Elem()))
	case reflect.Map:
		return fmt.Sprintf("map[%s]%s", file.Type(tpe.Key()), file.Type(tpe.Elem()))
	default:
		return tpe.String()
	}
}

func (file *File) Val(v interface{}) string {
	tpe := reflect.TypeOf(v)
	rv := reflect.ValueOf(v)

	switch rv.Kind() {
	case reflect.Map:
		parts := make([]string, 0)
		isMulti := rv.Len() > 1
		for _, key := range rv.MapKeys() {
			s := file.Val(key.Interface()) + ": " + file.Val(rv.MapIndex(key).Interface())
			if isMulti {
				parts = append(parts, s+",\n")
			} else {
				parts = append(parts, s)
			}
		}
		sort.Strings(parts)

		if isMulti {
			f := `%s{
				%s
			}`
			return fmt.Sprintf(f, file.Type(tpe), strings.Join(parts, ""))
		}
		f := "%s{%s}"
		return fmt.Sprintf(f, file.Type(tpe), strings.Join(parts, ", "))
	case reflect.Slice, reflect.Array:
		buf := new(bytes.Buffer)
		for i := 0; i < rv.Len(); i++ {
			s := file.Val(rv.Index(i).Interface())
			if i == 0 {
				buf.WriteString(s)
			} else {
				buf.WriteString(", " + s)
			}
		}
		return fmt.Sprintf("%s{%s}", file.Type(tpe), buf.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint8:
		return fmt.Sprintf("%d", v)
	case reflect.Bool:
		return strconv.FormatBool(v.(bool))
	case reflect.Float32:
		return strconv.FormatFloat(float64(v.(float32)), 'f', -1, 32)
	case reflect.Float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case reflect.Invalid:
		return "nil"
	case reflect.String:
		return strconv.Quote(v.(string))
	}
	return ""
}

func deVendor(importPath string) string {
	parts := strings.Split(importPath, "/vendor/")
	return parts[len(parts)-1]
}

package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitToWords(t *testing.T) {
	tt := assert.New(t)

	tt.Equal([]string{}, splitToWords(""))
	tt.Equal([]string{"lowercase"}, splitToWords("lowercase"))
	tt.Equal([]string{"Class"}, splitToWords("Class"))
	tt.Equal([]string{"My", "Class"}, splitToWords("MyClass"))
	tt.Equal([]string{"My", "C"}, splitToWords("MyC"))
	tt.Equal([]string{"HTML"}, splitToWords("HTML"))
	tt.Equal([]string{"PDF", "Loader"}, splitToWords("PDFLoader"))
	tt.Equal([]string{"A", "String"}, splitToWords("AString"))
	tt.Equal([]string{"Simple", "XML", "Parser"}, splitToWords("SimpleXMLParser"))
	tt.Equal([]string{"vim", "RPC", "Plugin"}, splitToWords("vimRPCPlugin"))
	tt.Equal([]string{"GL", "11", "Version"}, splitToWords("GL11Version"))
	tt.Equal([]string{"99", "Bottles"}, splitToWords("99Bottles"))
	tt.Equal([]string{"May", "5"}, splitToWords("May5"))
	tt.Equal([]string{"BFG", "9000"}, splitToWords("BFG9000"))
	tt.Equal([]string{"Böse", "Überraschung"}, splitToWords("BöseÜberraschung"))
	tt.Equal([]string{"Two", "spaces"}, splitToWords("Two  spaces"))
	tt.Equal([]string{"BadUTF8\xe2\xe2\xa1"}, splitToWords("BadUTF8\xe2\xe2\xa1"))
	tt.Equal([]string{"snake", "case"}, splitToWords("snake_case"))
	tt.Equal([]string{"snake", "case"}, splitToWords("snake_ case"))
}

func TestIdentifier(t *testing.T) {
	tt := assert.New(t)

	tt.Equal("SnakeCase", Identifier("snake_case").UpperCamelCase())
	tt.Equal("IDCase", Identifier("id_case").UpperCamelCase())

	tt.Equal("snakeCase", Identifier("snake_case").LowerCamelCase())
	tt.Equal("idCase", Identifier("id_case").LowerCamelCase())

	tt.Equal("SNAKE_CASE", Identifier("snakeCase").UpperSnakeCase())
	tt.Equal("ID_CASE", Identifier("idCase").UpperSnakeCase())

	tt.Equal("snake_case", Identifier("snakeCase").LowerSnakeCase())
	tt.Equal("id_case", Identifier("idCase").LowerSnakeCase())
	tt.Equal("i7_case", Identifier("i7Case").LowerSnakeCase())
}

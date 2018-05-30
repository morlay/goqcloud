package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFile_Sdump(t *testing.T) {
	tt := assert.New(t)
	file := NewFile("test")

	tt.Equal(`[]string{"1", "2"}`, file.Val([]string{"1", "2"}))
	tt.Equal(`[]interface {}{"1", nil}`, file.Val([]interface{}{"1", nil}))
	tt.Equal(`[2]interface {}{"1", nil}`, file.Val([2]interface{}{"1", nil}))
	tt.Equal(`map[string]int{"1": 2}`, file.Val(map[string]int{"1": 2}))
}

func TestFile(t *testing.T) {
	file := NewFile("test")

	file.WriteString("var a = 1")

	t.Log(string(file.Bytes()))
}

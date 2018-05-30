package generator

import (
	"os"
	"testing"

	"github.com/morlay/goqcloud/generator/codegen"
)

func init() {
	os.Chdir("..")
}

func TestAPIDoc(t *testing.T) {
	d := NewAPIDoc(&APIDocEntries{
		ContentsURL:  "https://cloud.tencent.com/document/api/213/15689",
		DataTypesURL: "https://cloud.tencent.com/document/api/213/15753",
	})

	d.Scan()

	f := codegen.NewFile(d.Service)

	d.APIs["DescribeInstances"].Write(f)

	t.Log(string(f.Bytes()))
}

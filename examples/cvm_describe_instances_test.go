package examples

import (
	"os"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"

	"github.com/morlay/goqcloud"
	"github.com/morlay/goqcloud/clients/cvm"
)

func TestCVMDescribeDescribeInstances(t *testing.T) {
	tt := assert.New(t)

	client := goqcloud.NewClientWithCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	req := cvm.DescribeInstancesRequest{
		Region: "ap-beijing",
	}
	resp, err := req.Invoke(client)
	tt.NoError(err)

	spew.Dump(resp.InstanceSet)

	tt.True(len(resp.InstanceSet) > 1)
}

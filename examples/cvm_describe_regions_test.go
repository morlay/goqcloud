package examples

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/morlay/goqcloud"
	"github.com/morlay/goqcloud/clients/cvm"
)

func TestCVMDescribeRegions(t *testing.T) {
	tt := assert.New(t)

	client := goqcloud.NewClientWithCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		goqcloud.ClientOptionWithTransports(goqcloud.NewLogTransport()),
	)

	req := cvm.DescribeRegionsRequest{}
	resp, err := req.Invoke(client)
	tt.NoError(err)

	tt.True(len(resp.RegionSet) > 1)

	for _, region := range resp.RegionSet {
		t.Log(region)
	}
}

func TestCVMDescribeRegionsWithError(t *testing.T) {
	tt := assert.New(t)

	client := goqcloud.NewClientWithCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		"",
	)

	req := cvm.DescribeRegionsRequest{}
	_, err := req.Invoke(client)
	tt.Error(err)
}

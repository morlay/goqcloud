package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 解挂云硬盘
// https://cloud.tencent.com/document/api/362/16316
type DetachDisksRequest struct {
	// 将要解挂的云硬盘ID， 通过DescribeDisks接口查询，单次请求最多可解挂10块弹性云盘。
	DiskIds []*string `name:"DiskIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DetachDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DetachDisksResponse, error) {
	resp := &DetachDisksResponse{}
	err := client.Request("cbs", "DetachDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type DetachDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

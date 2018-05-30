package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 挂载云硬盘
// https://cloud.tencent.com/document/api/362/16313
type AttachDisksRequest struct {
	// 将要被挂载的弹性云盘ID。通过DescribeDisks接口查询。单次最多可挂载10块弹性云盘。
	DiskIds []*string `name:"DiskIds"`
	// 云服务器实例ID。云盘将被挂载到此云服务器上，通过DescribeInstances接口查询。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *AttachDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AttachDisksResponse, error) {
	resp := &AttachDisksResponse{}
	err := client.Request("cbs", "AttachDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type AttachDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

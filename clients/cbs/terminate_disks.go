package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 退还云硬盘
// https://cloud.tencent.com/document/api/362/16321
type TerminateDisksRequest struct {
	// 需退还的云盘ID列表。
	DiskIds []*string `name:"DiskIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *TerminateDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TerminateDisksResponse, error) {
	resp := &TerminateDisksResponse{}
	err := client.Request("cbs", "TerminateDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type TerminateDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

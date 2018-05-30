package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 扩容云硬盘
// https://cloud.tencent.com/document/api/362/16310
type ResizeDiskRequest struct {
	// 云硬盘ID， 通过DescribeDisks接口查询。
	DiskId string `name:"DiskId"`
	// 云硬盘扩容后的大小，单位为GB，必须大于当前云硬盘大小。取值范围： 普通云硬盘:10GB ~ 4000G；高性能云硬盘:50GB ~ 4000GB；SSD云硬盘:100GB ~ 4000GB，步长均为10GB。
	DiskSize int64 `name:"DiskSize"`
	// 区域
	Region string `name:"Region"`
}

func (req *ResizeDiskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResizeDiskResponse, error) {
	resp := &ResizeDiskResponse{}
	err := client.Request("cbs", "ResizeDisk", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResizeDiskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

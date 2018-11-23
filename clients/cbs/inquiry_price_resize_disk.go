package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 扩容云硬盘询价
// https://cloud.tencent.com/document/api/362/16320

type InquiryPriceResizeDiskRequest struct {
	// 云硬盘ID， 通过DescribeDisks接口查询。
	DiskId string `name:"DiskId"`
	// 云硬盘扩容后的大小，单位为GB，不得小于当前云硬盘大小。云盘大小取值范围参见云硬盘产品分类的说明。
	DiskSize int64 `name:"DiskSize"`
	// 云盘所属项目ID。 如传入则仅用于鉴权。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *InquiryPriceResizeDiskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceResizeDiskResponse, error) {
	resp := &InquiryPriceResizeDiskResponse{}
	err := client.Request("cbs", "InquiryPriceResizeDisk", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceResizeDiskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 描述了扩容云盘的价格。
	DiskPrice PrepayPrice `json:"DiskPrice"`
}

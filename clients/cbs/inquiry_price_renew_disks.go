package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费云硬盘询价
// https://cloud.tencent.com/document/api/362/16317

type InquiryPriceRenewDisksRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的购买时长。如果在该参数中指定CurInstanceDeadline，则会按对齐到子机到期时间来续费。如果是批量续费询价，该参数与Disks参数一一对应，元素数量需保持一致。
	DiskChargePrepaids []*DiskChargePrepaid `name:"DiskChargePrepaids,omitempty"`
	// 云硬盘ID， 通过DescribeDisks接口查询。
	DiskIds []*string `name:"DiskIds"`
	// 指定云盘新的到期时间，形式如：2017-12-17 00:00:00。参数NewDeadline和DiskChargePrepaids是两种指定询价时长的方式，两者必传一个。
	NewDeadline *string `name:"NewDeadline,omitempty"`
	// 云盘所属项目ID。 如传入则仅用于鉴权。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *InquiryPriceRenewDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceRenewDisksResponse, error) {
	resp := &InquiryPriceRenewDisksResponse{}
	err := client.Request("cbs", "InquiryPriceRenewDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceRenewDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 描述了续费云盘的价格。
	DiskPrice Price `json:"DiskPrice"`
}

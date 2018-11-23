package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建云硬盘询价
// https://cloud.tencent.com/document/api/362/16314

type InquiryPriceCreateDisksRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `name:"DiskChargePrepaid,omitempty"`
	// 云硬盘计费类型。PREPAID：预付费，即包年包月POSTPAID_BY_HOUR：按小时后付费
	DiskChargeType string `name:"DiskChargeType"`
	// 购买云盘的数量。不填则默认为1。
	DiskCount *int64 `name:"DiskCount,omitempty"`
	// 云硬盘大小，单位为GB。云盘大小取值范围参见云硬盘产品分类的说明。
	DiskSize int64 `name:"DiskSize"`
	// 云硬盘类型。取值范围：普通云硬盘：CLOUD_BASIC高性能云硬盘：CLOUD_PREMIUMSSD云硬盘：CLOUD_SSD。
	DiskType string `name:"DiskType"`
	// 云盘所属项目ID。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *InquiryPriceCreateDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceCreateDisksResponse, error) {
	resp := &InquiryPriceCreateDisksResponse{}
	err := client.Request("cbs", "InquiryPriceCreateDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceCreateDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 描述了新购云盘的价格。
	DiskPrice Price `json:"DiskPrice"`
}

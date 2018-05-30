package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云硬盘配额
// https://cloud.tencent.com/document/api/362/16318
type DescribeDiskConfigQuotaRequest struct {
	// 实例CPU核数。
	CPU *int64 `name:"CPU,omitempty"`
	// 付费模式。取值范围：PREPAID：预付费POSTPAID_BY_HOUR：后付费。
	DiskChargeType *string `name:"DiskChargeType,omitempty"`
	// 硬盘介质类型。取值范围：CLOUD_BASIC：表示普通云硬盘CLOUD_PREMIUM：表示高性能云硬盘CLOUD_SSD：表示SSD云硬盘。
	DiskTypes []*string `name:"DiskTypes,omitempty"`
	// 系统盘或数据盘。取值范围：SYSTEM_DISK：表示系统盘DATA_DISK：表示数据盘。
	DiskUsage *string `name:"DiskUsage,omitempty"`
	// 查询类别，取值范围。INQUIRY_CBS_CONFIG：查询云盘配置列表INQUIRY_CVM_CONFIG：查询云盘与实例搭配的配置列表。
	InquiryType string `name:"InquiryType"`
	// 按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。详见实例类型
	InstanceFamilies []*string `name:"InstanceFamilies,omitempty"`
	// 实例内存大小。
	Memory *int64 `name:"Memory,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询一个或多个可用区下的配置。
	Zones []*string `name:"Zones,omitempty"`
}

func (req *DescribeDiskConfigQuotaRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDiskConfigQuotaResponse, error) {
	resp := &DescribeDiskConfigQuotaResponse{}
	err := client.Request("cbs", "DescribeDiskConfigQuota", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDiskConfigQuotaResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云盘配置列表。
	DiskConfigSet []*DiskConfig `json:"DiskConfigSet"`
}

package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取批量计算可用区机型配置信息
// https://cloud.tencent.com/document/api/599/18565

type DescribeCvmZoneInstanceConfigInfosRequest struct {
	// 过滤条件。 zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。 instance-family String - 是否必填：否 -（过滤条件）按照机型系列过滤。实例机型系列形如：S1、I1、M1等。 instance-type - String - 是否必填：否 - （过滤条件）按照机型过滤。 instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 ( POSTPAID_BY_HOUR：表示后付费，即按量计费机型 | SPOTPAID：表示竞价付费机型。 )
	Filters []*Filter `name:"Filters,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeCvmZoneInstanceConfigInfosRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCvmZoneInstanceConfigInfosResponse, error) {
	resp := &DescribeCvmZoneInstanceConfigInfosResponse{}
	err := client.Request("batch", "DescribeCvmZoneInstanceConfigInfos", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeCvmZoneInstanceConfigInfosResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 可用区机型配置列表。
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet"`
}

package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取可用区机型配置信息
// https://cloud.tencent.com/document/api/213/17378

type DescribeZoneInstanceConfigInfosRequest struct {
	// 过滤条件。 zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。 instance-family String - 是否必填：否 -（过滤条件）按照机型系列过滤。按照实例机型系列过滤。实例机型系列形如：S1、I1、M1等。 instance-type - String - 是否必填：否 - （过滤条件）按照机型过滤。按照实例机型过滤。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 DescribeInstanceTypeConfigs 来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。 instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示CDH付费，即只对CDH计费，不对CDH上的实例计费。 )
	Filters []*Filter `name:"Filters,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeZoneInstanceConfigInfosRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZoneInstanceConfigInfosResponse, error) {
	resp := &DescribeZoneInstanceConfigInfosResponse{}
	err := client.Request("cvm", "DescribeZoneInstanceConfigInfos", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeZoneInstanceConfigInfosResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 可用区机型配置列表。
	InstanceTypeQuotaSet []*InstanceTypeQuotaItem `json:"InstanceTypeQuotaSet"`
}

package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖价格
// https://cloud.tencent.com/document/api/409/16777

type InquiryPriceCreateDbInstancesRequest struct {
	// 实例计费类型。目前只支持：PREPAID（预付费，即包年包月）。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 实例数量。目前最大数量不超过100，如需一次性创建更多实例，请联系客服支持。
	InstanceCount int64 `name:"InstanceCount"`
	// 购买时长，单位：月。目前只支持1,2,3,4,5,6,7,8,9,10,11,12,24,36这些值。
	Period int64 `name:"Period"`
	// 计费ID。该参数可以通过调用DescribeProductConfig接口的返回值中的Pid字段来获取。
	Pid int64 `name:"Pid"`
	// 区域
	Region string `name:"Region"`
	// 规格ID。该参数可以通过调用DescribeProductConfig接口的返回值中的SpecCode字段来获取。
	SpecCode string `name:"SpecCode"`
	// 存储容量大小，单位：GB。
	Storage int64 `name:"Storage"`
	// 可用区ID。该参数可以通过调用 DescribeZones 接口的返回值中的Zone字段来获取。
	Zone string `name:"Zone"`
}

func (req *InquiryPriceCreateDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceCreateDbInstancesResponse, error) {
	resp := &InquiryPriceCreateDbInstancesResponse{}
	err := client.Request("postgres", "InquiryPriceCreateDBInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type InquiryPriceCreateDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原始价格，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 折后价格，单位：分
	Price int64 `json:"Price"`
}

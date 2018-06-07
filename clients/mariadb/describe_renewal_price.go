package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费实例询价
// https://cloud.tencent.com/document/api/237/16181

type DescribeRenewalPriceRequest struct {
	// 待续费的实例ID。形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRenewalPriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRenewalPriceResponse, error) {
	resp := &DescribeRenewalPriceResponse{}
	err := client.Request("mariadb", "DescribeRenewalPrice", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeRenewalPriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原价，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际价格，单位：分。受折扣等影响，可能和原价不同。
	Price int64 `json:"Price"`
}

package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费实例询价
// https://cloud.tencent.com/document/api/557/16138
type DescribeDcdbRenewalPriceRequest struct {
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 续费时长，单位：月。不传则默认为1个月。
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDcdbRenewalPriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDcdbRenewalPriceResponse, error) {
	resp := &DescribeDcdbRenewalPriceResponse{}
	err := client.Request("dcdb", "DescribeDCDBRenewalPrice", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDcdbRenewalPriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原价，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际价格，单位：分。受折扣等影响，可能和原价不同。
	Price int64 `json:"Price"`
}

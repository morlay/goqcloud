package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费分布式数据库实例
// https://cloud.tencent.com/document/api/557/16137

type RenewDcdbInstanceRequest struct {
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `name:"AutoVoucher,omitempty"`
	// 待续费的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 续费时长，单位：月。
	Period int64 `name:"Period"`
	// 区域
	Region string `name:"Region"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *RenewDcdbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewDcdbInstanceResponse, error) {
	resp := &RenewDcdbInstanceResponse{}
	err := client.Request("dcdb", "RenewDCDBInstance", "2018-04-11").Do(req, resp)
	return resp, err
}

type RenewDcdbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 长订单号。可以据此调用 DescribeOrders 查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName string `json:"DealName"`
}

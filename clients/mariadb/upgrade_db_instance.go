package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 扩容实例
// https://cloud.tencent.com/document/api/237/16189
type UpgradeDbInstanceRequest struct {
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `name:"AutoVoucher,omitempty"`
	// 待升级的实例ID。形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 内存大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得。
	Memory int64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 存储空间大小，单位：GB，可以通过 DescribeDBInstanceSpecs 查询实例规格获得不同内存大小对应的磁盘规格下限和上限。
	Storage int64 `name:"Storage"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds *string `name:"VoucherIds,omitempty"`
}

func (req *UpgradeDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeDbInstanceResponse, error) {
	resp := &UpgradeDbInstanceResponse{}
	err := client.Request("mariadb", "UpgradeDBInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type UpgradeDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 长订单号。可以据此调用 DescribeOrders 查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName string `json:"DealName"`
}

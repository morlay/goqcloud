package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 升级分布式数据库
// https://cloud.tencent.com/document/api/557/16136

type UpgradeDcdbInstanceRequest struct {
	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `name:"AddShardConfig,omitempty"`
	// 是否自动使用代金券进行支付，默认不使用。
	AutoVoucher *bool `name:"AutoVoucher,omitempty"`
	// 扩容分片配置，当UpgradeType为EXPAND时生效。
	ExpandShardConfig *ExpandShardConfig `name:"ExpandShardConfig,omitempty"`
	// 待升级的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 切分分片配置，当UpgradeType为SPLIT时生效。
	SplitShardConfig *SplitShardConfig `name:"SplitShardConfig,omitempty"`
	// 升级类型，取值范围:  ADD: 新增分片    EXPAND: 升级实例中的已有分片    SPLIT: 将已有分片中的数据切分到新增分片上
	UpgradeType string `name:"UpgradeType"`
	// 代金券ID列表，目前仅支持指定一张代金券。
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *UpgradeDcdbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeDcdbInstanceResponse, error) {
	resp := &UpgradeDcdbInstanceResponse{}
	err := client.Request("dcdb", "UpgradeDCDBInstance", "2018-04-11").Do(req, resp)
	return resp, err
}

type UpgradeDcdbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 长订单号。可以据此调用 DescribeOrders 查询订单详细信息，或在支付失败时调用用户账号相关接口进行支付。
	DealName string `json:"DealName"`
}

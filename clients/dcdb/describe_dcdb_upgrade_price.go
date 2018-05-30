package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询升级分布式数据库实例价格
// https://cloud.tencent.com/document/api/557/16132
type DescribeDcdbUpgradePriceRequest struct {
	// 新增分片配置，当UpgradeType为ADD时生效。
	AddShardConfig *AddShardConfig `name:"AddShardConfig,omitempty"`
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
}

func (req *DescribeDcdbUpgradePriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDcdbUpgradePriceResponse, error) {
	resp := &DescribeDcdbUpgradePriceResponse{}
	err := client.Request("dcdb", "DescribeDCDBUpgradePrice", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDcdbUpgradePriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原价，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际价格，单位：分。受折扣等影响，可能和原价不同。
	Price int64 `json:"Price"`
}

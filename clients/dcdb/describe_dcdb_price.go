package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新购分布式数据库实例询价
// https://cloud.tencent.com/document/api/557/16131
type DescribeDcdbPriceRequest struct {
	// 欲购买实例的数量，目前只支持购买1个实例
	Count int64 `name:"Count"`
	// 欲购买的时长，单位：月。
	Period int64 `name:"Period"`
	// 区域
	Region string `name:"Region"`
	// 实例分片个数，可选范围2-8，可以通过升级实例进行新增分片到最多64个分片。
	ShardCount int64 `name:"ShardCount"`
	// 分片内存大小，单位：GB，可以通过 DescribeShardSpec 查询实例规格获得。
	ShardMemory int64 `name:"ShardMemory"`
	// 单个分片节点个数大小，可以通过 DescribeShardSpec 查询实例规格获得。
	ShardNodeCount int64 `name:"ShardNodeCount"`
	// 分片存储空间大小，单位：GB，可以通过 DescribeShardSpec 查询实例规格获得。
	ShardStorage int64 `name:"ShardStorage"`
	// 欲新购实例的可用区ID。
	Zone string `name:"Zone"`
}

func (req *DescribeDcdbPriceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDcdbPriceResponse, error) {
	resp := &DescribeDcdbPriceResponse{}
	err := client.Request("dcdb", "DescribeDCDBPrice", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDcdbPriceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 原价，单位：分
	OriginalPrice int64 `json:"OriginalPrice"`
	// 实际价格，单位：分。受折扣等影响，可能和原价不同。
	Price int64 `json:"Price"`
}

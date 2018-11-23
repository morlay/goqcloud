package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询分片信息
// https://cloud.tencent.com/document/api/557/19990

type DescribeDcdbShardsRequest struct {
	// 实例ID，形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 返回数量，默认为 20，最大值为 100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为 0
	Offset *int64 `name:"Offset,omitempty"`
	// 排序字段， 目前仅支持 createtime
	OrderBy *string `name:"OrderBy,omitempty"`
	// 排序类型， desc 或者 asc
	OrderByType *string `name:"OrderByType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 分片Id列表。
	ShardInstanceIds []*string `name:"ShardInstanceIds,omitempty"`
}

func (req *DescribeDcdbShardsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDcdbShardsResponse, error) {
	resp := &DescribeDcdbShardsResponse{}
	err := client.Request("dcdb", "DescribeDCDBShards", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDcdbShardsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 分片信息列表
	Shards []*DCDBShardInfo `json:"Shards"`
	// 符合条件的分片数量
	TotalCount int64 `json:"TotalCount"`
}

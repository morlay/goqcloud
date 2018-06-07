package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取日志列表
// https://cloud.tencent.com/document/api/557/16133

type DescribeDbLogFilesRequest struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 分片 ID，形如：shard-7noic7tv
	ShardId string `name:"ShardId"`
	// 请求日志类型，取值只能为1、2、3或者4。1-binlog，2-冷备，3-errlog，4-slowlog。
	Type int64 `name:"Type"`
}

func (req *DescribeDbLogFilesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbLogFilesResponse, error) {
	resp := &DescribeDbLogFilesResponse{}
	err := client.Request("dcdb", "DescribeDBLogFiles", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDbLogFilesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

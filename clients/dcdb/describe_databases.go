package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库列表
// https://cloud.tencent.com/document/api/557/19987

type DescribeDatabasesRequest struct {
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDatabasesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDatabasesResponse, error) {
	resp := &DescribeDatabasesResponse{}
	err := client.Request("dcdb", "DescribeDatabases", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDatabasesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 该实例上的数据库列表。
	Databases []*Database `json:"Databases"`
	// 透传入参。
	InstanceId string `json:"InstanceId"`
}

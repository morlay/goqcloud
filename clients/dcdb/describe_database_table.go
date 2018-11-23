package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库表信息
// https://cloud.tencent.com/document/api/557/19988

type DescribeDatabaseTableRequest struct {
	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName string `name:"DbName"`
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 表名称，通过 DescribeDatabaseObjects 接口获取。
	Table string `name:"Table"`
}

func (req *DescribeDatabaseTableRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDatabaseTableResponse, error) {
	resp := &DescribeDatabaseTableResponse{}
	err := client.Request("dcdb", "DescribeDatabaseTable", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDatabaseTableResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 列信息。
	Cols []*TableColumn `json:"Cols"`
	// 数据库名称。
	DbName string `json:"DbName"`
	// 实例名称。
	InstanceId string `json:"InstanceId"`
	// 表名称。
	Table string `json:"Table"`
}

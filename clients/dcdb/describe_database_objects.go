package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询数据库对象列表
// https://cloud.tencent.com/document/api/557/19989

type DescribeDatabaseObjectsRequest struct {
	// 数据库名称，通过 DescribeDatabases 接口获取。
	DbName string `name:"DbName"`
	// 实例 ID，形如：dcdbt-ow7t8lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDatabaseObjectsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDatabaseObjectsResponse, error) {
	resp := &DescribeDatabaseObjectsResponse{}
	err := client.Request("dcdb", "DescribeDatabaseObjects", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDatabaseObjectsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据库名称。
	DbName string `json:"DbName"`
	// 函数列表。
	Funcs []*DatabaseFunction `json:"Funcs"`
	// 透传入参。
	InstanceId string `json:"InstanceId"`
	// 存储过程列表。
	Procs []*DatabaseProcedure `json:"Procs"`
	// 表列表。
	Tables []*DatabaseTable `json:"Tables"`
	// 视图列表。
	Views []*DatabaseView `json:"Views"`
}

package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 初始化实例
// https://cloud.tencent.com/document/api/237/16185

type InitDbInstancesRequest struct {
	// 待初始化的实例Id列表，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceIds []*string `name:"InstanceIds"`
	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。
	Params []*DBParamValue `name:"Params"`
	// 区域
	Region string `name:"Region"`
}

func (req *InitDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InitDbInstancesResponse, error) {
	resp := &InitDbInstancesResponse{}
	err := client.Request("mariadb", "InitDBInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type InitDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowId int64 `json:"FlowId"`
	// 透传入参。
	InstanceIds []*string `json:"InstanceIds"`
}

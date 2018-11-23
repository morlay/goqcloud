package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 初始化实例
// https://cloud.tencent.com/document/api/557/19985

type InitDcdbInstancesRequest struct {
	// 待初始化的实例Id列表，形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceIds []*string `name:"InstanceIds"`
	// 参数列表。本接口的可选值为：character_set_server（字符集，必传），lower_case_table_names（表名大小写敏感，必传），innodb_page_size（innodb数据页，默认16K），sync_mode（同步模式：0 - 异步； 1 - 强同步；2 - 强同步可退化。默认为强同步）。
	Params []*DBParamValue `name:"Params"`
	// 区域
	Region string `name:"Region"`
}

func (req *InitDcdbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InitDcdbInstancesResponse, error) {
	resp := &InitDcdbInstancesResponse{}
	err := client.Request("dcdb", "InitDCDBInstances", "2018-04-11").Do(req, resp)
	return resp, err
}

type InitDcdbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowIds []*int64 `json:"FlowIds"`
	// 透传入参。
	InstanceIds []*string `json:"InstanceIds"`
}

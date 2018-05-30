package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 初始化新实例
// https://cloud.tencent.com/document/api/236/15873
type InitDbInstancesRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值
	InstanceIds []*string `name:"InstanceIds"`
	// 实例新的密码，密码规则：8-64个字符，至少包含字母、数字、字符（支持的字符：!@#$%^*()）中的两种
	NewPassword string `name:"NewPassword"`
	// 实例的参数列表，目前支持设置“character_set_server”、“lower_case_table_names”参数。其中，“character_set_server”参数可选值为["utf8","latin1","gbk","utf8mb4"]；“lower_case_table_names”可选值为[“0”,“1”]
	Parameters []*ParamInfo `name:"Parameters"`
	// 区域
	Region string `name:"Region"`
	// 实例的端口
	Vport *int64 `name:"Vport,omitempty"`
}

func (req *InitDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InitDbInstancesResponse, error) {
	resp := &InitDbInstancesResponse{}
	err := client.Request("cdb", "InitDBInstances", "2017-03-20").Do(req, resp)
	return resp, err
}

type InitDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID数组，可使用此ID查询异步任务的执行结果
	AsyncRequestIds []*string `json:"AsyncRequestIds"`
}

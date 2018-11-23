package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重启实例
// https://cloud.tencent.com/document/api/238/19951

type RestartDbInstanceRequest struct {
	// 数据库实例ID，形如mssql-njj2mtpl
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RestartDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RestartDbInstanceResponse, error) {
	resp := &RestartDbInstanceResponse{}
	err := client.Request("sqlserver", "RestartDBInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type RestartDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务流程ID
	FlowId int64 `json:"FlowId"`
}

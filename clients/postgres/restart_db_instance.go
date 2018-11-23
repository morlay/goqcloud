package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重启实例
// https://cloud.tencent.com/document/api/409/18097

type RestartDbInstanceRequest struct {
	// 实例ID，形如postgres-6r233v55
	DBInstanceId string `name:"DBInstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RestartDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RestartDbInstanceResponse, error) {
	resp := &RestartDbInstanceResponse{}
	err := client.Request("postgres", "RestartDBInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type RestartDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步流程ID
	FlowId int64 `json:"FlowId"`
}

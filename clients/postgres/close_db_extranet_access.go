package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 关闭实例外网链接
// https://cloud.tencent.com/document/api/409/18105

type CloseDbExtranetAccessRequest struct {
	// 实例ID，形如postgres-6r233v55
	DBInstanceId string `name:"DBInstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CloseDbExtranetAccessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CloseDbExtranetAccessResponse, error) {
	resp := &CloseDbExtranetAccessResponse{}
	err := client.Request("postgres", "CloseDBExtranetAccess", "2017-03-12").Do(req, resp)
	return resp, err
}

type CloseDbExtranetAccessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务流程ID
	FlowId int64 `json:"FlowId"`
}

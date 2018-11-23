package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 开通外网
// https://cloud.tencent.com/document/api/409/18099

type OpenDbExtranetAccessRequest struct {
	// 实例ID，形如postgres-hez4fh0v
	DBInstanceId string `name:"DBInstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *OpenDbExtranetAccessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*OpenDbExtranetAccessResponse, error) {
	resp := &OpenDbExtranetAccessResponse{}
	err := client.Request("postgres", "OpenDBExtranetAccess", "2017-03-12").Do(req, resp)
	return resp, err
}

type OpenDbExtranetAccessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务流程ID
	FlowId int64 `json:"FlowId"`
}

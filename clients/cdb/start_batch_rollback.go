package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 回档数据库表
// https://cloud.tencent.com/document/api/236/18725

type StartBatchRollbackRequest struct {
	// 用于回档的实例详情信息
	Instances []*RollbackInstancesInfo `name:"Instances"`
	// 区域
	Region string `name:"Region"`
}

func (req *StartBatchRollbackRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StartBatchRollbackResponse, error) {
	resp := &StartBatchRollbackResponse{}
	err := client.Request("cdb", "StartBatchRollback", "2017-03-20").Do(req, resp)
	return resp, err
}

type StartBatchRollbackResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果
	AsyncRequestId string `json:"AsyncRequestId"`
}

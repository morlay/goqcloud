package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 终止数据导入任务
// https://cloud.tencent.com/document/api/236/15857

type StopDbImportJobRequest struct {
	// 异步任务的请求ID。
	AsyncRequestId string `name:"AsyncRequestId"`
	// 区域
	Region string `name:"Region"`
}

func (req *StopDbImportJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StopDbImportJobResponse, error) {
	resp := &StopDbImportJobResponse{}
	err := client.Request("cdb", "StopDBImportJob", "2017-03-20").Do(req, resp)
	return resp, err
}

type StopDbImportJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

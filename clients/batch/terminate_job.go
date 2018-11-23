package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 终止作业
// https://cloud.tencent.com/document/api/599/20409

type TerminateJobRequest struct {
	// 作业ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *TerminateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*TerminateJobResponse, error) {
	resp := &TerminateJobResponse{}
	err := client.Request("batch", "TerminateJob", "2017-03-12").Do(req, resp)
	return resp, err
}

type TerminateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

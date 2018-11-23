package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 开始灾备同步任务
// https://cloud.tencent.com/document/api/571/18568

type StartSyncJobRequest struct {
	// 灾备同步任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *StartSyncJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StartSyncJobResponse, error) {
	resp := &StartSyncJobResponse{}
	err := client.Request("dts", "StartSyncJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type StartSyncJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 校验灾备同步任务
// https://cloud.tencent.com/document/api/571/18574

type CreateSyncCheckJobRequest struct {
	// 灾备同步任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateSyncCheckJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSyncCheckJobResponse, error) {
	resp := &CreateSyncCheckJobResponse{}
	err := client.Request("dts", "CreateSyncCheckJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type CreateSyncCheckJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

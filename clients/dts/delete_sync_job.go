package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除灾备同步任务
// https://cloud.tencent.com/document/api/571/18572

type DeleteSyncJobRequest struct {
	// 待删除的灾备同步任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteSyncJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSyncJobResponse, error) {
	resp := &DeleteSyncJobResponse{}
	err := client.Request("dts", "DeleteSyncJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type DeleteSyncJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

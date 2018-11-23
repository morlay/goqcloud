package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 撤销数据迁移任务
// https://cloud.tencent.com/document/api/571/18136

type StopMigrateJobRequest struct {
	// 数据迁移任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *StopMigrateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StopMigrateJobResponse, error) {
	resp := &StopMigrateJobResponse{}
	err := client.Request("dts", "StopMigrateJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type StopMigrateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

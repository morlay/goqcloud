package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 启动数据迁移任务
// https://cloud.tencent.com/document/api/571/18137

type StartMigrateJobRequest struct {
	// 数据迁移任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *StartMigrateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StartMigrateJobResponse, error) {
	resp := &StartMigrateJobResponse{}
	err := client.Request("dts", "StartMigrateJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type StartMigrateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

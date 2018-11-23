package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 完成数据迁移任务
// https://cloud.tencent.com/document/api/571/18143

type CompleteMigrateJobRequest struct {
	// 数据迁移任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CompleteMigrateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CompleteMigrateJobResponse, error) {
	resp := &CompleteMigrateJobResponse{}
	err := client.Request("dts", "CompleteMigrateJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type CompleteMigrateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

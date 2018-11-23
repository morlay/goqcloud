package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建校验迁移任务
// https://cloud.tencent.com/document/api/571/18142

type CreateMigrateCheckJobRequest struct {
	// 数据迁移任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateMigrateCheckJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateMigrateCheckJobResponse, error) {
	resp := &CreateMigrateCheckJobResponse{}
	err := client.Request("dts", "CreateMigrateCheckJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type CreateMigrateCheckJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

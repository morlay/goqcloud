package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除数据迁移任务
// https://cloud.tencent.com/document/api/571/18140

type DeleteMigrateJobRequest struct {
	// 数据迁移任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteMigrateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteMigrateJobResponse, error) {
	resp := &DeleteMigrateJobResponse{}
	err := client.Request("dts", "DeleteMigrateJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type DeleteMigrateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

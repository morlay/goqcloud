package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除作业
// https://cloud.tencent.com/document/api/599/15906

type DeleteJobRequest struct {
	// 作业ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteJobResponse, error) {
	resp := &DeleteJobResponse{}
	err := client.Request("batch", "DeleteJob", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

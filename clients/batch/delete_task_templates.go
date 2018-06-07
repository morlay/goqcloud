package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除任务模板
// https://cloud.tencent.com/document/api/599/15900

type DeleteTaskTemplatesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 用于删除任务模板信息
	TaskTemplateIds []*string `name:"TaskTemplateIds"`
}

func (req *DeleteTaskTemplatesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteTaskTemplatesResponse, error) {
	resp := &DeleteTaskTemplatesResponse{}
	err := client.Request("batch", "DeleteTaskTemplates", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteTaskTemplatesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

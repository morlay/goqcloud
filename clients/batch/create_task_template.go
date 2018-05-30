package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建任务模板
// https://cloud.tencent.com/document/api/599/15899
type CreateTaskTemplateRequest struct {
	// 区域
	Region string `name:"Region"`
	// 任务模板描述
	TaskTemplateDescription *string `name:"TaskTemplateDescription,omitempty"`
	// 任务模板内容，参数要求与任务一致
	TaskTemplateInfo Task `name:"TaskTemplateInfo"`
	// 任务模板名称
	TaskTemplateName string `name:"TaskTemplateName"`
}

func (req *CreateTaskTemplateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateTaskTemplateResponse, error) {
	resp := &CreateTaskTemplateResponse{}
	err := client.Request("batch", "CreateTaskTemplate", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateTaskTemplateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务模板ID
	TaskTemplateId string `json:"TaskTemplateId"`
}

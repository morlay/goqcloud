package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改任务模板
// https://cloud.tencent.com/document/api/599/15901

type ModifyTaskTemplateRequest struct {
	// 区域
	Region string `name:"Region"`
	// 任务模板描述
	TaskTemplateDescription *string `name:"TaskTemplateDescription,omitempty"`
	// 任务模板ID
	TaskTemplateId string `name:"TaskTemplateId"`
	// 任务模板信息
	TaskTemplateInfo *Task `name:"TaskTemplateInfo,omitempty"`
	// 任务模板名称
	TaskTemplateName *string `name:"TaskTemplateName,omitempty"`
}

func (req *ModifyTaskTemplateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyTaskTemplateResponse, error) {
	resp := &ModifyTaskTemplateResponse{}
	err := client.Request("batch", "ModifyTaskTemplate", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyTaskTemplateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

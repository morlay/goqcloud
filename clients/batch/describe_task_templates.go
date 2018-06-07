package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取任务模板
// https://cloud.tencent.com/document/api/599/15902

type DescribeTaskTemplatesRequest struct {
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 任务模板ID
	TaskTemplateIds []*string `name:"TaskTemplateIds,omitempty"`
}

func (req *DescribeTaskTemplatesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskTemplatesResponse, error) {
	resp := &DescribeTaskTemplatesResponse{}
	err := client.Request("batch", "DescribeTaskTemplates", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeTaskTemplatesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务模板列表
	TaskTemplateSet []*TaskTemplateView `json:"TaskTemplateSet"`
	// 任务模板数量
	TotalCount int64 `json:"TotalCount"`
}

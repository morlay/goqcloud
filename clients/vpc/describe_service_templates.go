package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询协议端口模板
// https://cloud.tencent.com/document/api/215/16719
type DescribeServiceTemplatesRequest struct {
	// 过滤条件。service-template-name - String - （过滤条件）协议端口模板名称。service-template-id - String - （过滤条件）协议端口模板实例ID，例如：ppm-e6dy460g。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。
	Limit *string `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *string `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeServiceTemplatesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeServiceTemplatesResponse, error) {
	resp := &DescribeServiceTemplatesResponse{}
	err := client.Request("vpc", "DescribeServiceTemplates", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeServiceTemplatesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 协议端口模板对象。
	ServiceTemplateSet []*ServiceTemplate `json:"ServiceTemplateSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}

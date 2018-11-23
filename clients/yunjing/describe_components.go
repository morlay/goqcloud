package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取组件列表
// https://cloud.tencent.com/document/api/296/30336

type DescribeComponentsRequest struct {
	// 组件ID。Uuid和ComponentId必填其一，使用ComponentId表示，查询该组件列表信息。
	ComponentId *int64 `name:"ComponentId,omitempty"`
	// 过滤条件。ComponentVersion - String - 是否必填：否 - 组件版本号MachineIp - String - 是否必填：否 - 主机内网IP
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。Uuid和ComponentId必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeComponentsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeComponentsResponse, error) {
	resp := &DescribeComponentsResponse{}
	err := client.Request("yunjing", "DescribeComponents", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeComponentsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 组件列表数据。
	Components []*Component `json:"Components"`
	// 组件列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}

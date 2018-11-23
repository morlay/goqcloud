package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取安全事件消息
// https://cloud.tencent.com/document/api/296/30330

type DescribeSecurityDynamicsRequest struct {
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSecurityDynamicsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSecurityDynamicsResponse, error) {
	resp := &DescribeSecurityDynamicsResponse{}
	err := client.Request("yunjing", "DescribeSecurityDynamics", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeSecurityDynamicsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全事件消息数组。
	SecurityDynamics []*SecurityDynamic `json:"SecurityDynamics"`
	// 记录总数。
	TotalCount int64 `json:"TotalCount"`
}

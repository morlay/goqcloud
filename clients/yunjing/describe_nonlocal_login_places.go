package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取异地登录事件
// https://cloud.tencent.com/document/api/296/19836

type DescribeNonlocalLoginPlacesRequest struct {
	// 过滤条件。Keywords - String - 是否必填：否 -  查询关键字Status - String - 是否必填：否 -  登录状态（NON_LOCAL_LOGIN: 异地登录 | NORMAL_LOGIN : 正常登录）
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 客户端唯一Uuid。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeNonlocalLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeNonlocalLoginPlacesResponse, error) {
	resp := &DescribeNonlocalLoginPlacesResponse{}
	err := client.Request("yunjing", "DescribeNonlocalLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeNonlocalLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异地登录信息数组。
	NonLocalLoginPlaces []*NonLocalLoginPlace `json:"NonLocalLoginPlaces"`
	// 记录总数。
	TotalCount int64 `json:"TotalCount"`
}

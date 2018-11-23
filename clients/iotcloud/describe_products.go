package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取产品列表
// https://cloud.tencent.com/document/api/634/19477

type DescribeProductsRequest struct {
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 分页大小，当前页面中显示的最大数量，值范围 10-250。
	Limit int64 `name:"Limit"`
	// 分页偏移，Offset从0开始
	Offset int64 `name:"Offset"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeProductsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProductsResponse, error) {
	resp := &DescribeProductsResponse{}
	err := client.Request("iotcloud", "DescribeProducts", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeProductsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 产品详细信息列表
	Products []*ProductInfo `json:"Products"`
	// 产品总数
	TotalCount int64 `json:"TotalCount"`
}

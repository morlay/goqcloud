package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询地域列表
// https://cloud.tencent.com/document/api/213/15708

type DescribeRegionsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRegionsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("cvm", "DescribeRegions", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 地域列表信息
	RegionSet []*RegionInfo `json:"RegionSet"`
	// 地域数量
	TotalCount int64 `json:"TotalCount"`
}

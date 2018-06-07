package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖地域
// https://cloud.tencent.com/document/api/409/16768

type DescribeRegionsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRegionsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("postgres", "DescribeRegions", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 地域信息集合。
	RegionSet []*RegionInfo `json:"RegionSet"`
	// 返回的结果数量。
	TotalCount int64 `json:"TotalCount"`
}

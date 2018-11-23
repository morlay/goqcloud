package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖地域
// https://cloud.tencent.com/document/api/238/19965

type DescribeRegionsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRegionsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRegionsResponse, error) {
	resp := &DescribeRegionsResponse{}
	err := client.Request("sqlserver", "DescribeRegions", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeRegionsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 地域信息数组
	RegionSet []*RegionInfo `json:"RegionSet"`
	// 返回地域信息总的条目
	TotalCount int64 `json:"TotalCount"`
}

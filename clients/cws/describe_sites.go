package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看站点列表
// https://cloud.tencent.com/document/api/692/16752
type DescribeSitesRequest struct {
	// 过滤条件
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 站点ID列表
	SiteIds []*int64 `name:"SiteIds,omitempty"`
}

func (req *DescribeSitesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSitesResponse, error) {
	resp := &DescribeSitesResponse{}
	err := client.Request("cws", "DescribeSites", "2018-03-12").Do(req, resp)
	return resp, err
}

type DescribeSitesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 站点信息列表。
	Sites []*Site `json:"Sites"`
	// 站点数量。
	TotalCount int64 `json:"TotalCount"`
}

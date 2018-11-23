package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取专业周报异地登录数据
// https://cloud.tencent.com/document/api/296/30323

type DescribeWeeklyReportNonlocalLoginPlacesRequest struct {
	// 专业周报开始时间。
	BeginDate time.Time `name:"BeginDate"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeWeeklyReportNonlocalLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeWeeklyReportNonlocalLoginPlacesResponse, error) {
	resp := &DescribeWeeklyReportNonlocalLoginPlacesResponse{}
	err := client.Request("yunjing", "DescribeWeeklyReportNonlocalLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeWeeklyReportNonlocalLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 记录总数。
	TotalCount int64 `json:"TotalCount"`
	// 专业周报异地登录数据。
	WeeklyReportNonlocalLoginPlaces []*WeeklyReportNonlocalLoginPlace `json:"WeeklyReportNonlocalLoginPlaces"`
}

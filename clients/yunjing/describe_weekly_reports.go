package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取周报列表
// https://cloud.tencent.com/document/api/296/30321

type DescribeWeeklyReportsRequest struct {
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeWeeklyReportsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeWeeklyReportsResponse, error) {
	resp := &DescribeWeeklyReportsResponse{}
	err := client.Request("yunjing", "DescribeWeeklyReports", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeWeeklyReportsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 记录总数。
	TotalCount int64 `json:"TotalCount"`
	// 专业周报列表数组。
	WeeklyReports []*WeeklyReport `json:"WeeklyReports"`
}

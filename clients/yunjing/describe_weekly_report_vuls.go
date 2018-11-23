package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取专业版周报漏洞数据
// https://cloud.tencent.com/document/api/296/30322

type DescribeWeeklyReportVulsRequest struct {
	// 专业版周报开始时间。
	BeginDate time.Time `name:"BeginDate"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeWeeklyReportVulsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeWeeklyReportVulsResponse, error) {
	resp := &DescribeWeeklyReportVulsResponse{}
	err := client.Request("yunjing", "DescribeWeeklyReportVuls", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeWeeklyReportVulsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 记录总数。
	TotalCount int64 `json:"TotalCount"`
	// 专业周报漏洞数据数组。
	WeeklyReportVuls []*WeeklyReportVul `json:"WeeklyReportVuls"`
}

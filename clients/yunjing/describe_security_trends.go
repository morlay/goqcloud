package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取安全事件统计数据
// https://cloud.tencent.com/document/api/296/30329

type DescribeSecurityTrendsRequest struct {
	// 开始时间。
	BeginDate time.Time `name:"BeginDate"`
	// 结束时间。
	EndDate time.Time `name:"EndDate"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSecurityTrendsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSecurityTrendsResponse, error) {
	resp := &DescribeSecurityTrendsResponse{}
	err := client.Request("yunjing", "DescribeSecurityTrends", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeSecurityTrendsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 基线统计数据数组。
	BaseLines []*SecurityTrend `json:"BaseLines"`
	// 密码破解事件统计数据数组。
	BruteAttacks []*SecurityTrend `json:"BruteAttacks"`
	// 木马事件统计数据数组。
	Malwares []*SecurityTrend `json:"Malwares"`
	// 异地登录事件统计数据数组。
	NonLocalLoginPlaces []*SecurityTrend `json:"NonLocalLoginPlaces"`
	// 漏洞统计数据数组。
	Vuls []*SecurityTrend `json:"Vuls"`
}

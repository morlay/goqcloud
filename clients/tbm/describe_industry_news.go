package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取定制任务媒体报道列表
// https://cloud.tencent.com/document/api/853/18401

type DescribeIndustryNewsRequest struct {
	// 查询结束时间
	EndDate time.Time `name:"EndDate"`
	// 行业ID
	IndustryId string `name:"IndustryId"`
	// 查询条数上限，默认20
	Limit *int64 `name:"Limit,omitempty"`
	// 查询偏移，默认从0开始
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 是否显示列表，若为 true，则返回文章列表
	ShowList *bool `name:"ShowList,omitempty"`
	// 查询开始时间
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeIndustryNewsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeIndustryNewsResponse, error) {
	resp := &DescribeIndustryNewsResponse{}
	err := client.Request("tbm", "DescribeIndustryNews", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeIndustryNewsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 总计疑似负面数量
	AdverseCount int64 `json:"AdverseCount"`
	// 按天统计的数量列表
	DateCountSet []*DateCount `json:"DateCountSet"`
	// 总计来源数量
	FromCount int64 `json:"FromCount"`
	// 总计文章数量
	NewsCount int64 `json:"NewsCount"`
	// 文章列表
	NewsSet []*IndustryNews `json:"NewsSet"`
}

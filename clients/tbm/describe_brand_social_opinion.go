package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取品牌社交渠道观点
// https://cloud.tencent.com/document/api/853/18394

type DescribeBrandSocialOpinionRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 检索结束时间
	EndDate time.Time `name:"EndDate"`
	// 查询条数上限，默认20
	Limit *int64 `name:"Limit,omitempty"`
	// 查询偏移，默认从0开始
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 列表显示标记，若为true，则返回文章列表详情
	ShowList *bool `name:"ShowList,omitempty"`
	// 检索开始时间
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeBrandSocialOpinionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBrandSocialOpinionResponse, error) {
	resp := &DescribeBrandSocialOpinionResponse{}
	err := client.Request("tbm", "DescribeBrandSocialOpinion", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeBrandSocialOpinionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 疑似负面报道总数
	AdverseCount int64 `json:"AdverseCount"`
	// 文章总数
	ArticleCount int64 `json:"ArticleCount"`
	// 文章列表详情
	ArticleSet []*BrandReportArticle `json:"ArticleSet"`
	// 来源统计总数
	FromCount int64 `json:"FromCount"`
}

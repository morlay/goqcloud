package tbm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取品牌媒体报道数
// https://cloud.tencent.com/document/api/853/18397

type DescribeBrandMediaReportRequest struct {
	// 品牌ID
	BrandId string `name:"BrandId"`
	// 查询结束时间
	EndDate time.Time `name:"EndDate"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartDate time.Time `name:"StartDate"`
}

func (req *DescribeBrandMediaReportRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBrandMediaReportResponse, error) {
	resp := &DescribeBrandMediaReportResponse{}
	err := client.Request("tbm", "DescribeBrandMediaReport", "2018-01-29").Do(req, resp)
	return resp, err
}

type DescribeBrandMediaReportResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 按天计算的每天文章数
	DateCountSet []*DateCount `json:"DateCountSet"`
	// 查询范围内文章总数
	TotalCount int64 `json:"TotalCount"`
}

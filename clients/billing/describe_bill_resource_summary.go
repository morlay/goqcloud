package billing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询账单资源汇总数据
// https://cloud.tencent.com/document/api/555/19181

type DescribeBillResourceSummaryRequest struct {
	// 数量，最大值为1000
	Limit int64 `name:"Limit"`
	// 月份，格式为yyyy-mm
	Month string `name:"Month"`
	// 偏移量
	Offset int64 `name:"Offset"`
	// 周期类型，byUsedTime按计费周期/byPayTime按扣费周期
	PeriodType string `name:"PeriodType"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBillResourceSummaryRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBillResourceSummaryResponse, error) {
	resp := &DescribeBillResourceSummaryResponse{}
	err := client.Request("billing", "DescribeBillResourceSummary", "2018-07-09").Do(req, resp)
	return resp, err
}

type DescribeBillResourceSummaryResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 资源汇总列表
	ResourceSummarySet []*BillResourceSummary `json:"ResourceSummarySet"`
}

package billing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询账单明细数据
// https://cloud.tencent.com/document/api/555/19182

type DescribeBillDetailRequest struct {
	// 数量，最大值为100
	Limit int64 `name:"Limit"`
	// 月份，格式为yyyy-mm
	Month string `name:"Month"`
	// 偏移量
	Offset int64 `name:"Offset"`
	// 周期类型，byPayTime按扣费周期/byUsedTime按计费周期
	PeriodType string `name:"PeriodType"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBillDetailRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBillDetailResponse, error) {
	resp := &DescribeBillDetailResponse{}
	err := client.Request("billing", "DescribeBillDetail", "2018-07-09").Do(req, resp)
	return resp, err
}

type DescribeBillDetailResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 详情列表
	DetailSet []*BillDetail `json:"DetailSet"`
}

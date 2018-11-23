package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取预授权规则列表
// https://cloud.tencent.com/document/api/386/30141

type DescribePsaRegulationsRequest struct {
	// 数量限制
	Limit int64 `name:"Limit"`
	// 偏移量
	Offset int64 `name:"Offset"`
	// 排序方式 0:递增(默认) 1:递减
	Order *int64 `name:"Order,omitempty"`
	// 排序字段，取值支持：CreateTime
	OrderField *string `name:"OrderField,omitempty"`
	// 规则ID过滤，支持模糊查询
	PsaIds []*string `name:"PsaIds,omitempty"`
	// 规则别名过滤，支持模糊查询
	PsaNames []*string `name:"PsaNames,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 标签过滤
	Tags []*Tag `name:"Tags,omitempty"`
}

func (req *DescribePsaRegulationsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePsaRegulationsResponse, error) {
	resp := &DescribePsaRegulationsResponse{}
	err := client.Request("bm", "DescribePsaRegulations", "2018-04-23").Do(req, resp)
	return resp, err
}

type DescribePsaRegulationsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回规则列表
	PsaRegulations []*PsaRegulation `json:"PsaRegulations"`
	// 返回规则数量
	TotalCount int64 `json:"TotalCount"`
}

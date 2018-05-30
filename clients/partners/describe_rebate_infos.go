package partners

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询代理商返佣信息
// https://cloud.tencent.com/document/api/563/16041
type DescribeRebateInfosRequest struct {
	// 限制数目
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量
	Offset *int64 `name:"Offset,omitempty"`
	// 返佣月份，如2018-02
	RebateMonth *string `name:"RebateMonth,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRebateInfosRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRebateInfosResponse, error) {
	resp := &DescribeRebateInfosResponse{}
	err := client.Request("partners", "DescribeRebateInfos", "2018-03-21").Do(req, resp)
	return resp, err
}

type DescribeRebateInfosResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返佣信息列表
	RebateInfoSet []*RebateInfoElem `json:"RebateInfoSet"`
	// 符合查询条件返佣信息数目
	TotalCount int64 `json:"TotalCount"`
}

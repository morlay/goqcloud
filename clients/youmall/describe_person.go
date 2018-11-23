package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询客户信息
// https://cloud.tencent.com/document/api/860/30166

type DescribePersonRequest struct {
	// 查询数量，默认20，最大查询数量100
	Limit *int64 `name:"Limit,omitempty"`
	// 卖场编码
	MallId string `name:"MallId"`
	// 查询偏移
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribePersonRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePersonResponse, error) {
	resp := &DescribePersonResponse{}
	err := client.Request("youmall", "DescribePerson", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribePersonResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 客户信息
	PersonSet []*PersonProfile `json:"PersonSet"`
	// 总计客户数量
	TotalCount int64 `json:"TotalCount"`
}

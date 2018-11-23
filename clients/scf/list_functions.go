package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取函数列表
// https://cloud.tencent.com/document/api/583/18582

type ListFunctionsRequest struct {
	// 返回数据长度，默认值为 20
	Limit *int64 `name:"Limit,omitempty"`
	// 数据偏移量，默认值为 0
	Offset *int64 `name:"Offset,omitempty"`
	// 以升序还是降序的方式返回结果，可选值 ASC 和 DESC
	Order *string `name:"Order,omitempty"`
	// 根据哪个字段进行返回结果排序,支持以下字段：AddTime, ModTime, FunctionName
	Orderby *string `name:"Orderby,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 支持FunctionName模糊匹配
	SearchKey *string `name:"SearchKey,omitempty"`
}

func (req *ListFunctionsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ListFunctionsResponse, error) {
	resp := &ListFunctionsResponse{}
	err := client.Request("scf", "ListFunctions", "2018-04-16").Do(req, resp)
	return resp, err
}

type ListFunctionsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 函数列表
	Functions []*Function `json:"Functions"`
	// 总数
	TotalCount int64 `json:"TotalCount"`
}

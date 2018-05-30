package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取产品列表
// https://cloud.tencent.com/document/api/568/16450
type GetProductsRequest struct {
	// 长度
	Length *int64 `name:"Length,omitempty"`
	// 偏移
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetProductsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetProductsResponse, error) {
	resp := &GetProductsResponse{}
	err := client.Request("iot", "GetProducts", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetProductsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Product列表
	Products []*Product `json:"Products"`
	// Product总数
	Total int64 `json:"Total"`
}

package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取产品信息
// https://cloud.tencent.com/document/api/568/16380

type GetProductRequest struct {
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetProductRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetProductResponse, error) {
	resp := &GetProductResponse{}
	err := client.Request("iot", "GetProduct", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetProductResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 产品信息
	Product Product `json:"Product"`
}

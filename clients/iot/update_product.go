package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新产品信息
// https://cloud.tencent.com/document/api/568/16451

type UpdateProductRequest struct {
	// 数据模版（json）
	DataTemplate *string `name:"DataTemplate,omitempty"`
	// 产品描述
	Description *string `name:"Description,omitempty"`
	// 产品名称
	Name *string `name:"Name,omitempty"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *UpdateProductRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateProductResponse, error) {
	resp := &UpdateProductResponse{}
	err := client.Request("iot", "UpdateProduct", "2018-01-23").Do(req, resp)
	return resp, err
}

type UpdateProductResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 更新后的产品信息
	Product Product `json:"Product"`
}

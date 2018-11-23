package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建产品
// https://cloud.tencent.com/document/api/634/19479

type CreateProductRequest struct {
	// 产品名称，名称不能和已经存在的产品名称重复。命名规则：[a-zA-Z0-9:_-]{1,32}
	ProductName string `name:"ProductName"`
	// 产品属性
	ProductProperties *ProductProperties `name:"ProductProperties,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateProductRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateProductResponse, error) {
	resp := &CreateProductResponse{}
	err := client.Request("iotcloud", "CreateProduct", "2018-06-14").Do(req, resp)
	return resp, err
}

type CreateProductResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 产品 ID，腾讯云生成全局唯一 ID
	ProductId string `json:"ProductId"`
	// 产品名称
	ProductName string `json:"ProductName"`
	// 产品属性
	ProductProperties ProductProperties `json:"ProductProperties"`
}

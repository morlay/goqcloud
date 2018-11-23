package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除产品
// https://cloud.tencent.com/document/api/634/19478

type DeleteProductRequest struct {
	// 需要删除的产品 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteProductRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteProductResponse, error) {
	resp := &DeleteProductResponse{}
	err := client.Request("iotcloud", "DeleteProduct", "2018-06-14").Do(req, resp)
	return resp, err
}

type DeleteProductResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

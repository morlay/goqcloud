package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除产品
// https://cloud.tencent.com/document/api/568/16449

type DeleteProductRequest struct {
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteProductRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteProductResponse, error) {
	resp := &DeleteProductResponse{}
	err := client.Request("iot", "DeleteProduct", "2018-01-23").Do(req, resp)
	return resp, err
}

type DeleteProductResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

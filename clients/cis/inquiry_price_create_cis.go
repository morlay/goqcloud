package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建容器实例询价
// https://cloud.tencent.com/document/api/858/17769

type InquiryPriceCreateCisRequest struct {
	// CPU，单位：核
	Cpu float64 `name:"Cpu"`
	// 内存，单位：Gi
	Memory float64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 可用区
	Zone string `name:"Zone"`
}

func (req *InquiryPriceCreateCisRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InquiryPriceCreateCisResponse, error) {
	resp := &InquiryPriceCreateCisResponse{}
	err := client.Request("cis", "InquiryPriceCreateCis", "2018-04-08").Do(req, resp)
	return resp, err
}

type InquiryPriceCreateCisResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 价格
	Price Price `json:"Price"`
}

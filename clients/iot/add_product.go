package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增产品
// https://cloud.tencent.com/document/api/568/16448
type AddProductRequest struct {
	// 鉴权模式（1：动态令牌，推荐使用动态令牌）
	AuthType *int64 `name:"AuthType,omitempty"`
	// 数据协议（native表示自定义，template表示数据模板，默认值为template）
	DataProtocol *string `name:"DataProtocol,omitempty"`
	// 数据模版（json数组）
	DataTemplate []*string `name:"DataTemplate,omitempty"`
	// 产品描述
	Description string `name:"Description"`
	// 产品名称，同一区域产品名称需唯一，支持中文、英文字母、中划线和下划线，长度不超过31个字符，中文占两个字符
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
}

func (req *AddProductRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddProductResponse, error) {
	resp := &AddProductResponse{}
	err := client.Request("iot", "AddProduct", "2018-01-23").Do(req, resp)
	return resp, err
}

type AddProductResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 产品信息
	Product Product `json:"Product"`
}

package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取Topic列表
// https://cloud.tencent.com/document/api/568/16447

type GetTopicsRequest struct {
	// 长度
	Length *int64 `name:"Length,omitempty"`
	// 偏移
	Offset *int64 `name:"Offset,omitempty"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *GetTopicsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetTopicsResponse, error) {
	resp := &GetTopicsResponse{}
	err := client.Request("iot", "GetTopics", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetTopicsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Topic列表
	Topics []*Topic `json:"Topics"`
	// Topic总数
	Total int64 `json:"Total"`
}

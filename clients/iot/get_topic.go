package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取Topic信息
// https://cloud.tencent.com/document/api/568/16446

type GetTopicRequest struct {
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// TopicId
	TopicId string `name:"TopicId"`
}

func (req *GetTopicRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetTopicResponse, error) {
	resp := &GetTopicResponse{}
	err := client.Request("iot", "GetTopic", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetTopicResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Topic信息
	Topic Topic `json:"Topic"`
}

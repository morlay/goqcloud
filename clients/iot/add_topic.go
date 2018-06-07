package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增Topic
// https://cloud.tencent.com/document/api/568/16444

type AddTopicRequest struct {
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// Topic名称
	TopicName string `name:"TopicName"`
}

func (req *AddTopicRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddTopicResponse, error) {
	resp := &AddTopicResponse{}
	err := client.Request("iot", "AddTopic", "2018-01-23").Do(req, resp)
	return resp, err
}

type AddTopicResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Topic信息
	Topic Topic `json:"Topic"`
}

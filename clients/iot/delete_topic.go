package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除Topic
// https://cloud.tencent.com/document/api/568/16445

type DeleteTopicRequest struct {
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// TopicId
	TopicId string `name:"TopicId"`
}

func (req *DeleteTopicRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteTopicResponse, error) {
	resp := &DeleteTopicResponse{}
	err := client.Request("iot", "DeleteTopic", "2018-01-23").Do(req, resp)
	return resp, err
}

type DeleteTopicResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

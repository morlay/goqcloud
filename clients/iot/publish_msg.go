package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 向Topic发布消息
// https://cloud.tencent.com/document/api/568/16453
type PublishMsgRequest struct {
	// 消息内容
	Message string `name:"Message"`
	// Qos(目前QoS支持0与1)
	Qos *int64 `name:"Qos,omitempty"`
	// 区域
	Region string `name:"Region"`
	// Topic
	Topic string `name:"Topic"`
}

func (req *PublishMsgRequest) Invoke(client github_com_morlay_goqcloud.Client) (*PublishMsgResponse, error) {
	resp := &PublishMsgResponse{}
	err := client.Request("iot", "PublishMsg", "2018-01-23").Do(req, resp)
	return resp, err
}

type PublishMsgResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

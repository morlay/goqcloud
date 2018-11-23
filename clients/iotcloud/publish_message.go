package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 发布消息
// https://cloud.tencent.com/document/api/634/19486

type PublishMessageRequest struct {
	// 设备名称
	DeviceName string `name:"DeviceName"`
	// 消息内容
	Payload string `name:"Payload"`
	// 产品ID
	ProductId string `name:"ProductId"`
	// 服务质量等级，取值为0， 1
	Qos *int64 `name:"Qos,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 消息发往的主题。命名规则：${ProductId}/${DeviceName}/[a-zA-Z0-9:_-]{1,128}
	Topic string `name:"Topic"`
}

func (req *PublishMessageRequest) Invoke(client github_com_morlay_goqcloud.Client) (*PublishMessageResponse, error) {
	resp := &PublishMessageResponse{}
	err := client.Request("iotcloud", "PublishMessage", "2018-06-14").Do(req, resp)
	return resp, err
}

type PublishMessageResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

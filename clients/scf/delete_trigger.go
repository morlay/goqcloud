package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除触发器
// https://cloud.tencent.com/document/api/583/18588

type DeleteTriggerRequest struct {
	// 函数的名称
	FunctionName string `name:"FunctionName"`
	// 函数的版本信息
	Qualifier *string `name:"Qualifier,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 如果删除的触发器类型为 COS 触发器，该字段为必填值，存放 JSON 格式的数据 {"event":"cos:ObjectCreated:*"}，数据内容和 SetTrigger 接口中该字段的格式相同；如果删除的触发器类型为定时触发器或 CMQ 触发器，可以不指定该字段
	TriggerDesc *string `name:"TriggerDesc,omitempty"`
	// 要删除的触发器名称
	TriggerName string `name:"TriggerName"`
	// 要删除的触发器类型，目前支持 cos 、cmq、 timer、ckafka 类型
	Type string `name:"Type"`
}

func (req *DeleteTriggerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteTriggerResponse, error) {
	resp := &DeleteTriggerResponse{}
	err := client.Request("scf", "DeleteTrigger", "2018-04-16").Do(req, resp)
	return resp, err
}

type DeleteTriggerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

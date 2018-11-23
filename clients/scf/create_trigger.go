package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置函数触发方式
// https://cloud.tencent.com/document/api/583/18589

type CreateTriggerRequest struct {
	// 新建触发器绑定的函数名称
	FunctionName string `name:"FunctionName"`
	// 函数的版本
	Qualifier *string `name:"Qualifier,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 触发器对应的参数，如果是 timer 类型的触发器其内容是 Linux cron 表达式，如果是其他触发器，见具体触发器说明
	TriggerDesc *string `name:"TriggerDesc,omitempty"`
	// 新建触发器名称。如果是定时触发器，名称支持英文字母、数字、连接符和下划线，最长100个字符；如果是其他触发器，见具体触发器绑定参数的说明
	TriggerName string `name:"TriggerName"`
	// 触发器类型，目前支持 cos 、cmq、 timers、 ckafka类型
	Type string `name:"Type"`
}

func (req *CreateTriggerRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateTriggerResponse, error) {
	resp := &CreateTriggerResponse{}
	err := client.Request("scf", "CreateTrigger", "2018-04-16").Do(req, resp)
	return resp, err
}

type CreateTriggerResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

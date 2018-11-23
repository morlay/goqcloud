package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 运行函数
// https://cloud.tencent.com/document/api/583/17243

type InvokeRequest struct {
	// 运行函数时的参数，以json格式传入，最大支持的参数长度是 1M
	ClientContext *string `name:"ClientContext,omitempty"`
	// 函数名称
	FunctionName string `name:"FunctionName"`
	// RequestResponse(同步) 和 Event(异步)，默认为同步
	InvocationType *string `name:"InvocationType,omitempty"`
	// 同步调用时指定该字段，返回值会包含4K的日志，可选值为None和Tail，默认值为None。当该值为Tail时，返回参数中的logMsg字段会包含对应的函数执行日志
	LogType *string `name:"LogType,omitempty"`
	// 命名空间
	Namespace *string `name:"Namespace,omitempty"`
	// 触发函数的版本号
	Qualifier *string `name:"Qualifier,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *InvokeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InvokeResponse, error) {
	resp := &InvokeResponse{}
	err := client.Request("scf", "Invoke", "2018-04-16").Do(req, resp)
	return resp, err
}

type InvokeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 函数执行结果
	Result Result `json:"Result"`
}

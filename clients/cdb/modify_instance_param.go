package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例参数
// https://cloud.tencent.com/document/api/236/15860
type ModifyInstanceParamRequest struct {
	// 实例短Id列表。
	InstanceIds []*string `name:"InstanceIds"`
	// 要修改的参数列表。每一个元素是name和currentValue的组合。name是参数名，currentValue是要修改成的值。
	ParamList []*Parameter `name:"ParamList"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyInstanceParamRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyInstanceParamResponse, error) {
	resp := &ModifyInstanceParamResponse{}
	err := client.Request("cdb", "ModifyInstanceParam", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyInstanceParamResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可用于查询任务进度。
	AsyncRequestId string `json:"AsyncRequestId"`
}

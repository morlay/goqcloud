package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库参数
// https://cloud.tencent.com/document/api/557/19982

type ModifyDbParametersRequest struct {
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `name:"Params"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbParametersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbParametersResponse, error) {
	resp := &ModifyDbParametersResponse{}
	err := client.Request("dcdb", "ModifyDBParameters", "2018-04-11").Do(req, resp)
	return resp, err
}

type ModifyDbParametersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId string `json:"InstanceId"`
	// 各参数修改结果
	Result []*ParamModifyResult `json:"Result"`
}

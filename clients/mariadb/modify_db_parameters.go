package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库参数
// https://cloud.tencent.com/document/api/237/16153
type ModifyDbParametersRequest struct {
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 参数列表，每一个元素是Param和Value的组合
	Params []*DBParamValue `name:"Params"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDbParametersRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbParametersResponse, error) {
	resp := &ModifyDbParametersResponse{}
	err := client.Request("mariadb", "ModifyDBParameters", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDbParametersResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 参数修改结果
	Config []*ParamModifyResult `json:"Config"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `json:"InstanceId"`
}

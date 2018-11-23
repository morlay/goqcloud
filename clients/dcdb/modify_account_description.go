package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库账号备注
// https://cloud.tencent.com/document/api/557/19984

type ModifyAccountDescriptionRequest struct {
	// 新的账号备注，长度 0~256。
	Description string `name:"Description"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host string `name:"Host"`
	// 实例 ID，形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 登录用户名。
	UserName string `name:"UserName"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountDescriptionResponse, error) {
	resp := &ModifyAccountDescriptionResponse{}
	err := client.Request("dcdb", "ModifyAccountDescription", "2018-04-11").Do(req, resp)
	return resp, err
}

type ModifyAccountDescriptionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

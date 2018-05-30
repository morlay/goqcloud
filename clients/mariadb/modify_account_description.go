package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库账号备注
// https://cloud.tencent.com/document/api/237/16170
type ModifyAccountDescriptionRequest struct {
	// 新的账号备注，长度 0~256。
	Description string `name:"Description"`
	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host string `name:"Host"`
	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 登录用户名。
	UserName string `name:"UserName"`
}

func (req *ModifyAccountDescriptionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAccountDescriptionResponse, error) {
	resp := &ModifyAccountDescriptionResponse{}
	err := client.Request("mariadb", "ModifyAccountDescription", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyAccountDescriptionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

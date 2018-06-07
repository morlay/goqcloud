package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置账号密码
// https://cloud.tencent.com/document/api/237/16168

type ResetAccountPasswordRequest struct {
	// 用户允许的访问 host，用户名+host唯一确定一个账号。
	Host string `name:"Host"`
	// 实例 ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 新密码，由字母、数字或常见符号组成，不能包含分号、单引号和双引号，长度为6~32位。
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
	// 登录用户名。
	UserName string `name:"UserName"`
}

func (req *ResetAccountPasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetAccountPasswordResponse, error) {
	resp := &ResetAccountPasswordResponse{}
	err := client.Request("mariadb", "ResetAccountPassword", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetAccountPasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

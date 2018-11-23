package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 重置账户密码
// https://cloud.tencent.com/document/api/409/18107

type ResetAccountPasswordRequest struct {
	// 实例ID，形如postgres-4wdeb0zv
	DBInstanceId string `name:"DBInstanceId"`
	// UserName账户对应的新密码
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
	// 实例账户名
	UserName string `name:"UserName"`
}

func (req *ResetAccountPasswordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResetAccountPasswordResponse, error) {
	resp := &ResetAccountPasswordResponse{}
	err := client.Request("postgres", "ResetAccountPassword", "2017-03-12").Do(req, resp)
	return resp, err
}

type ResetAccountPasswordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

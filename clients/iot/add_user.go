package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 注册用户
// https://cloud.tencent.com/document/api/568/17133

type AddUserRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *AddUserRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddUserResponse, error) {
	resp := &AddUserResponse{}
	err := client.Request("iot", "AddUser", "2018-01-23").Do(req, resp)
	return resp, err
}

type AddUserResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 用户信息
	User User `json:"User"`
}

package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 注册应用用户
// https://cloud.tencent.com/document/api/568/17134
type AppAddUserRequest struct {
	// 密码
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
	// 用户名
	UserName string `name:"UserName"`
}

func (req *AppAddUserRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AppAddUserResponse, error) {
	resp := &AppAddUserResponse{}
	err := client.Request("iot", "AppAddUser", "2018-01-23").Do(req, resp)
	return resp, err
}

type AppAddUserResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 应用用户
	AppUser AppUser `json:"AppUser"`
}

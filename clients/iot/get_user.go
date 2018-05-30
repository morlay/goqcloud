package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取用户信息
// https://cloud.tencent.com/document/api/568/16589
type GetUserRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *GetUserRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetUserResponse, error) {
	resp := &GetUserResponse{}
	err := client.Request("iot", "GetUser", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetUserResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 用户信息
	User User `json:"User"`
}

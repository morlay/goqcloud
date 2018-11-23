package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建自定义脚本
// https://cloud.tencent.com/document/api/386/30258

type CreateUserCmdRequest struct {
	// 用户自定义脚本的名称
	Alias string `name:"Alias"`
	// 脚本内容，必须经过base64编码
	Content string `name:"Content"`
	// 命令适用的操作系统类型，取值linux或xserver
	OsType string `name:"OsType"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateUserCmdRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateUserCmdResponse, error) {
	resp := &CreateUserCmdResponse{}
	err := client.Request("bm", "CreateUserCmd", "2018-04-23").Do(req, resp)
	return resp, err
}

type CreateUserCmdResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 脚本ID
	CmdId string `json:"CmdId"`
}

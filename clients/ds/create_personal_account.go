package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 个人开户
// https://cloud.tencent.com/document/api/869/17793

type CreatePersonalAccountRequest struct {
	// 个人用户证件号码
	IdentNo string `name:"IdentNo"`
	// 个人用户证件类型。0代表身份证
	IdentType int64 `name:"IdentType"`
	// 个人用户手机号
	MobilePhone string `name:"MobilePhone"`
	// 模块名
	Module string `name:"Module"`
	// 个人用户姓名
	Name string `name:"Name"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreatePersonalAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreatePersonalAccountResponse, error) {
	resp := &CreatePersonalAccountResponse{}
	err := client.Request("ds", "CreatePersonalAccount", "2018-05-23").Do(req, resp)
	return resp, err
}

type CreatePersonalAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 账号ID
	AccountResId string `json:"AccountResId"`
}

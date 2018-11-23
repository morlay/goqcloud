package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 检测验证码
// https://cloud.tencent.com/document/api/869/17796

type CheckVcodeRequest struct {
	// 帐号ID
	AccountResId string `name:"AccountResId"`
	// 合同ID
	ContractResId string `name:"ContractResId"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
	// 验证码
	VerifyCode string `name:"VerifyCode"`
}

func (req *CheckVcodeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CheckVcodeResponse, error) {
	resp := &CheckVcodeResponse{}
	err := client.Request("ds", "CheckVcode", "2018-05-23").Do(req, resp)
	return resp, err
}

type CheckVcodeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

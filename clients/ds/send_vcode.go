package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 发送验证码
// https://cloud.tencent.com/document/api/869/17787

type SendVcodeRequest struct {
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
}

func (req *SendVcodeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SendVcodeResponse, error) {
	resp := &SendVcodeResponse{}
	err := client.Request("ds", "SendVcode", "2018-05-23").Do(req, resp)
	return resp, err
}

type SendVcodeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

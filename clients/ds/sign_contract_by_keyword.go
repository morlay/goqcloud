package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 按关键字签署合同
// https://cloud.tencent.com/document/api/869/17960

type SignContractByKeywordRequest struct {
	// 账户ID
	AccountResId string `name:"AccountResId"`
	// 授权时间，格式为年月日时分秒，例20160801095509
	AuthorizationTime string `name:"AuthorizationTime"`
	// 合同ID
	ContractResId string `name:"ContractResId"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 授权IP地址
	Position string `name:"Position"`
	// 区域
	Region string `name:"Region"`
	// 签章ID
	SealResId *string `name:"SealResId,omitempty"`
	// 签署关键字，坐标和范围不得超过合同文件边界
	SignKeyword SignKeyword `name:"SignKeyword"`
}

func (req *SignContractByKeywordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SignContractByKeywordResponse, error) {
	resp := &SignContractByKeywordResponse{}
	err := client.Request("ds", "SignContractByKeyword", "2018-05-23").Do(req, resp)
	return resp, err
}

type SignContractByKeywordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

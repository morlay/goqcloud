package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 按坐标签署合同
// https://cloud.tencent.com/document/api/869/17786

type SignContractByCoordinateRequest struct {
	// 帐户ID
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
	// 印章ID
	SealResId *string `name:"SealResId,omitempty"`
	// 签署坐标，坐标不得超过合同文件边界
	SignLocations []*SignLocation `name:"SignLocations"`
}

func (req *SignContractByCoordinateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SignContractByCoordinateResponse, error) {
	resp := &SignContractByCoordinateResponse{}
	err := client.Request("ds", "SignContractByCoordinate", "2018-05-23").Do(req, resp)
	return resp, err
}

type SignContractByCoordinateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

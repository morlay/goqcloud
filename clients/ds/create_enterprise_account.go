package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 企业开户
// https://cloud.tencent.com/document/api/869/17794

type CreateEnterpriseAccountRequest struct {
	// 企业用户营业执照号码
	IdentNo string `name:"IdentNo"`
	// 企业用户证件类型，8代表营业执照
	IdentType int64 `name:"IdentType"`
	// 企业联系电话
	MobilePhone string `name:"MobilePhone"`
	// 模块名
	Module string `name:"Module"`
	// 企业用户名称
	Name string `name:"Name"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
	// 经办人证件号码
	TransactorIdentNo string `name:"TransactorIdentNo"`
	// 经办人证件类型，0代表身份证
	TransactorIdentType int64 `name:"TransactorIdentType"`
	// 经办人姓名
	TransactorName string `name:"TransactorName"`
	// 经办人手机号
	TransactorPhone string `name:"TransactorPhone"`
}

func (req *CreateEnterpriseAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateEnterpriseAccountResponse, error) {
	resp := &CreateEnterpriseAccountResponse{}
	err := client.Request("ds", "CreateEnterpriseAccount", "2018-05-23").Do(req, resp)
	return resp, err
}

type CreateEnterpriseAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 帐号ID
	AccountResId string `json:"AccountResId"`
}

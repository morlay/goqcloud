package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建集团门店管理员账号
// https://cloud.tencent.com/document/api/860/20421

type CreateAccountRequest struct {
	// 集团ID
	CompanyId string `name:"CompanyId"`
	// 账号名；需要是手机号
	Name string `name:"Name"`
	// 密码；需要是(`~!@#$%^&*()_+=-）中的至少两种且八位以上
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
	// 备注说明; 30个字符以内
	Remark *string `name:"Remark,omitempty"`
	// 客户门店编码
	ShopCode string `name:"ShopCode"`
}

func (req *CreateAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAccountResponse, error) {
	resp := &CreateAccountResponse{}
	err := client.Request("youmall", "CreateAccount", "2018-02-28").Do(req, resp)
	return resp, err
}

type CreateAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

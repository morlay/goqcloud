package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除帐号
// https://cloud.tencent.com/document/api/869/17791

type DeleteAccountRequest struct {
	// 帐号ID列表
	AccountList []*string `name:"AccountList"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteAccountRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteAccountResponse, error) {
	resp := &DeleteAccountResponse{}
	err := client.Request("ds", "DeleteAccount", "2018-05-23").Do(req, resp)
	return resp, err
}

type DeleteAccountResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 删除失败帐号ID列表
	DelFailedList []*string `json:"DelFailedList"`
	// 删除成功帐号ID列表
	DelSuccessList []*string `json:"DelSuccessList"`
}

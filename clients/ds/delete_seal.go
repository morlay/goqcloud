package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除签章
// https://cloud.tencent.com/document/api/869/17790

type DeleteSealRequest struct {
	// 帐号ID
	AccountResId string `name:"AccountResId"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
	// 签章ID
	SealResId string `name:"SealResId"`
}

func (req *DeleteSealRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSealResponse, error) {
	resp := &DeleteSealResponse{}
	err := client.Request("ds", "DeleteSeal", "2018-05-23").Do(req, resp)
	return resp, err
}

type DeleteSealResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 签章ID
	SealResId string `json:"SealResId"`
}

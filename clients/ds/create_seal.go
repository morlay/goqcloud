package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增签章
// https://cloud.tencent.com/document/api/869/17792

type CreateSealRequest struct {
	// 帐号ID
	AccountResId string `name:"AccountResId"`
	// 签章链接，图片必须为png格式
	ImgUrl string `name:"ImgUrl"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateSealRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSealResponse, error) {
	resp := &CreateSealResponse{}
	err := client.Request("ds", "CreateSeal", "2018-05-23").Do(req, resp)
	return resp, err
}

type CreateSealResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 签章ID
	SealResId string `json:"SealResId"`
}

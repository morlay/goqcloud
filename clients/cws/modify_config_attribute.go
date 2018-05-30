package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改用户配置的属性
// https://cloud.tencent.com/document/api/692/16758
type ModifyConfigAttributeRequest struct {
	// 漏洞告警通知等级，4位分别代表：高危、中危、低危、提示
	NoticeLevel *string `name:"NoticeLevel,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyConfigAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyConfigAttributeResponse, error) {
	resp := &ModifyConfigAttributeResponse{}
	err := client.Request("cws", "ModifyConfigAttribute", "2018-03-12").Do(req, resp)
	return resp, err
}

type ModifyConfigAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

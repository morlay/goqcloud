package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改站点的属性
// https://cloud.tencent.com/document/api/692/16754

type ModifySiteAttributeRequest struct {
	// 站点名称
	Name *string `name:"Name,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 站点ID
	SiteId int64 `name:"SiteId"`
}

func (req *ModifySiteAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifySiteAttributeResponse, error) {
	resp := &ModifySiteAttributeResponse{}
	err := client.Request("cws", "ModifySiteAttribute", "2018-03-12").Do(req, resp)
	return resp, err
}

type ModifySiteAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

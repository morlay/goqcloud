package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改站点的属性
// https://cloud.tencent.com/document/api/692/16754

type ModifySiteAttributeRequest struct {
	// 用于测试cookie是否有效的关键字
	LoginCheckKw *string `name:"LoginCheckKw,omitempty"`
	// 用于测试cookie是否有效的URL
	LoginCheckUrl *string `name:"LoginCheckUrl,omitempty"`
	// 登录后的cookie
	LoginCookie *string `name:"LoginCookie,omitempty"`
	// 站点名称
	Name *string `name:"Name,omitempty"`
	// 网站是否需要登录扫描：0-未知；-1-不需要；1-需要
	NeedLogin *int64 `name:"NeedLogin,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 禁止扫描器扫描的目录关键字
	ScanDisallow *string `name:"ScanDisallow,omitempty"`
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

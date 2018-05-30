package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增站点扫描任务
// https://cloud.tencent.com/document/api/692/16749
type CreateSitesScansRequest struct {
	// 扫描速率限制，每秒发送X个HTTP请求
	RateLimit int64 `name:"RateLimit"`
	// 区域
	Region string `name:"Region"`
	// 扫描模式，normal-正常扫描；deep-深度扫描
	ScannerType string `name:"ScannerType"`
	// 站点的ID列表
	SiteIds []*int64 `name:"SiteIds"`
}

func (req *CreateSitesScansRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSitesScansResponse, error) {
	resp := &CreateSitesScansResponse{}
	err := client.Request("cws", "CreateSitesScans", "2018-03-12").Do(req, resp)
	return resp, err
}

type CreateSitesScansResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

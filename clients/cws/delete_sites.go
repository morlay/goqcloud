package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除站点
// https://cloud.tencent.com/document/api/692/16750
type DeleteSitesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 站点ID列表
	SiteIds []*int64 `name:"SiteIds"`
}

func (req *DeleteSitesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSitesResponse, error) {
	resp := &DeleteSitesResponse{}
	err := client.Request("cws", "DeleteSites", "2018-03-12").Do(req, resp)
	return resp, err
}

type DeleteSitesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

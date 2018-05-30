package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增站点
// https://cloud.tencent.com/document/api/692/16748
type CreateSitesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 站点的url列表
	Urls []*string `name:"Urls"`
}

func (req *CreateSitesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSitesResponse, error) {
	resp := &CreateSitesResponse{}
	err := client.Request("cws", "CreateSites", "2018-03-12").Do(req, resp)
	return resp, err
}

type CreateSitesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 新增站点数。
	Number int64 `json:"Number"`
}

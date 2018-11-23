package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 忽略漏洞
// https://cloud.tencent.com/document/api/296/19856

type IgnoreImpactedHostsRequest struct {
	// 漏洞ID数组。
	Ids []*int64 `name:"Ids"`
	// 区域
	Region string `name:"Region"`
}

func (req *IgnoreImpactedHostsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*IgnoreImpactedHostsResponse, error) {
	resp := &IgnoreImpactedHostsResponse{}
	err := client.Request("yunjing", "IgnoreImpactedHosts", "2018-02-28").Do(req, resp)
	return resp, err
}

type IgnoreImpactedHostsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

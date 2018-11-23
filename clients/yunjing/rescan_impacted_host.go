package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 漏洞重新检测
// https://cloud.tencent.com/document/api/296/19855

type RescanImpactedHostRequest struct {
	// 漏洞ID。
	Id int64 `name:"Id"`
	// 区域
	Region string `name:"Region"`
}

func (req *RescanImpactedHostRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RescanImpactedHostResponse, error) {
	resp := &RescanImpactedHostResponse{}
	err := client.Request("yunjing", "RescanImpactedHost", "2018-02-28").Do(req, resp)
	return resp, err
}

type RescanImpactedHostResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

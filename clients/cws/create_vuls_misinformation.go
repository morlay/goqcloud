package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增漏洞误报信息
// https://cloud.tencent.com/document/api/692/16740
type CreateVulsMisinformationRequest struct {
	// 区域
	Region string `name:"Region"`
	// 漏洞ID列表
	VulIds []*int64 `name:"VulIds"`
}

func (req *CreateVulsMisinformationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateVulsMisinformationResponse, error) {
	resp := &CreateVulsMisinformationResponse{}
	err := client.Request("cws", "CreateVulsMisinformation", "2018-03-12").Do(req, resp)
	return resp, err
}

type CreateVulsMisinformationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

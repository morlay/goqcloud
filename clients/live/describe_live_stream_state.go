package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询流状态
// https://cloud.tencent.com/document/api/267/20470

type DescribeLiveStreamStateRequest struct {
	// 应用名称。
	AppName string `name:"AppName"`
	// 您的加速域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
	// 流名称。
	StreamName string `name:"StreamName"`
}

func (req *DescribeLiveStreamStateRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLiveStreamStateResponse, error) {
	resp := &DescribeLiveStreamStateResponse{}
	err := client.Request("live", "DescribeLiveStreamState", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamStateResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 流状态
	StreamState string `json:"StreamState"`
}

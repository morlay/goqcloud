package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 恢复延播
// https://cloud.tencent.com/document/api/267/20464

type ResumeDelayLiveStreamRequest struct {
	// 应用名称。
	AppName string `name:"AppName"`
	// 您的加速域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
	// 流名称。
	StreamName string `name:"StreamName"`
}

func (req *ResumeDelayLiveStreamRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResumeDelayLiveStreamResponse, error) {
	resp := &ResumeDelayLiveStreamResponse{}
	err := client.Request("live", "ResumeDelayLiveStream", "2018-08-01").Do(req, resp)
	return resp, err
}

type ResumeDelayLiveStreamResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

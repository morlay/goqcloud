package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 禁播直播流
// https://cloud.tencent.com/document/api/267/20468

type ForbidLiveStreamRequest struct {
	// 应用名称。
	AppName string `name:"AppName"`
	// 您的加速域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
	// 恢复流的时间。UTC 格式，例如：2018-11-29T19:00:00Z。注意：默认禁播90天，且最长支持禁播90天。
	ResumeTime *string `name:"ResumeTime,omitempty"`
	// 流名称。
	StreamName string `name:"StreamName"`
}

func (req *ForbidLiveStreamRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ForbidLiveStreamResponse, error) {
	resp := &ForbidLiveStreamResponse{}
	err := client.Request("live", "ForbidLiveStream", "2018-08-01").Do(req, resp)
	return resp, err
}

type ForbidLiveStreamResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 恢复直播流
// https://cloud.tencent.com/document/api/267/20467

type ResumeLiveStreamRequest struct {
	// 应用名称。
	AppName string `name:"AppName"`
	// 您的加速域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
	// 流名称。
	StreamName string `name:"StreamName"`
}

func (req *ResumeLiveStreamRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ResumeLiveStreamResponse, error) {
	resp := &ResumeLiveStreamResponse{}
	err := client.Request("live", "ResumeLiveStream", "2018-08-01").Do(req, resp)
	return resp, err
}

type ResumeLiveStreamResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

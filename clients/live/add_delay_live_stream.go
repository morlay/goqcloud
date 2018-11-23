package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 延迟播放
// https://cloud.tencent.com/document/api/267/20465

type AddDelayLiveStreamRequest struct {
	// 应用名称。
	AppName string `name:"AppName"`
	// 延播时间，单位：秒，上限：600秒。
	DelayTime int64 `name:"DelayTime"`
	// 您的加速域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
	// 流名称。
	StreamName string `name:"StreamName"`
}

func (req *AddDelayLiveStreamRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddDelayLiveStreamResponse, error) {
	resp := &AddDelayLiveStreamResponse{}
	err := client.Request("live", "AddDelayLiveStream", "2018-08-01").Do(req, resp)
	return resp, err
}

type AddDelayLiveStreamResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

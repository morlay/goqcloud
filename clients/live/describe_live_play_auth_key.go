package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询播放鉴权key
// https://cloud.tencent.com/document/api/267/30426

type DescribeLivePlayAuthKeyRequest struct {
	// 域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeLivePlayAuthKeyRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLivePlayAuthKeyResponse, error) {
	resp := &DescribeLivePlayAuthKeyResponse{}
	err := client.Request("live", "DescribeLivePlayAuthKey", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLivePlayAuthKeyResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 播放鉴权key信息。
	PlayAuthKeyInfo PlayAuthKeyInfo `json:"PlayAuthKeyInfo"`
}

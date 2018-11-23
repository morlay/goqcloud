package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询推流鉴权key
// https://cloud.tencent.com/document/api/267/30425

type DescribeLivePushAuthKeyRequest struct {
	// 推流域名。
	DomainName string `name:"DomainName"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeLivePushAuthKeyRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLivePushAuthKeyResponse, error) {
	resp := &DescribeLivePushAuthKeyResponse{}
	err := client.Request("live", "DescribeLivePushAuthKey", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLivePushAuthKeyResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 推流鉴权key信息。
	PushAuthKeyInfo PushAuthKeyInfo `json:"PushAuthKeyInfo"`
}

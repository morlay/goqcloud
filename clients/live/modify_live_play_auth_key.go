package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改播放鉴权key
// https://cloud.tencent.com/document/api/267/30424

type ModifyLivePlayAuthKeyRequest struct {
	// 有效时间，单位：秒。
	AuthDelta *int64 `name:"AuthDelta,omitempty"`
	// 鉴权key。
	AuthKey *string `name:"AuthKey,omitempty"`
	// 域名。
	DomainName string `name:"DomainName"`
	// 是否启用，0：关闭，1：启用。
	Enable *int64 `name:"Enable,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyLivePlayAuthKeyRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyLivePlayAuthKeyResponse, error) {
	resp := &ModifyLivePlayAuthKeyResponse{}
	err := client.Request("live", "ModifyLivePlayAuthKey", "2018-08-01").Do(req, resp)
	return resp, err
}

type ModifyLivePlayAuthKeyResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

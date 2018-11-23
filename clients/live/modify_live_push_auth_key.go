package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改推流鉴权key
// https://cloud.tencent.com/document/api/267/30423

type ModifyLivePushAuthKeyRequest struct {
	// 有效时间，单位：秒。
	AuthDelta *int64 `name:"AuthDelta,omitempty"`
	// 备鉴权key。
	BackupAuthKey *string `name:"BackupAuthKey,omitempty"`
	// 推流域名。
	DomainName string `name:"DomainName"`
	// 是否启用，0：关闭，1：启用。
	Enable *int64 `name:"Enable,omitempty"`
	// 主鉴权key。
	MasterAuthKey *string `name:"MasterAuthKey,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyLivePushAuthKeyRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyLivePushAuthKeyResponse, error) {
	resp := &ModifyLivePushAuthKeyResponse{}
	err := client.Request("live", "ModifyLivePushAuthKey", "2018-08-01").Do(req, resp)
	return resp, err
}

type ModifyLivePushAuthKeyResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

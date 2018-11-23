package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置自动续费
// https://cloud.tencent.com/document/api/409/18096

type SetAutoRenewFlagRequest struct {
	// 续费标记。0-正常续费；1-自动续费；2-到期不续费
	AutoRenewFlag int64 `name:"AutoRenewFlag"`
	// 实例ID数组
	DBInstanceIdSet []*string `name:"DBInstanceIdSet"`
	// 区域
	Region string `name:"Region"`
}

func (req *SetAutoRenewFlagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SetAutoRenewFlagResponse, error) {
	resp := &SetAutoRenewFlagResponse{}
	err := client.Request("postgres", "SetAutoRenewFlag", "2017-03-12").Do(req, resp)
	return resp, err
}

type SetAutoRenewFlagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设置成功的实例个数
	Count int64 `json:"Count"`
}

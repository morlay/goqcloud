package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费实例
// https://cloud.tencent.com/document/api/238/19953

type RenewDbInstanceRequest struct {
	// 是否自动使用代金券，0-不使用；1-使用；默认不实用
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 实例ID，形如mssql-j8kv137v
	InstanceId string `name:"InstanceId"`
	// 续费多少个月，取值范围为1-48，默认为1
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 代金券ID数组，目前只支持使用1张代金券
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *RenewDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewDbInstanceResponse, error) {
	resp := &RenewDbInstanceResponse{}
	err := client.Request("sqlserver", "RenewDBInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type RenewDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单名称
	DealName string `json:"DealName"`
}

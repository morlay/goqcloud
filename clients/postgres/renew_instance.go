package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费实例
// https://cloud.tencent.com/document/api/409/18098

type RenewInstanceRequest struct {
	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 实例ID，形如postgres-6fego161
	DBInstanceId string `name:"DBInstanceId"`
	// 续费多少个月
	Period int64 `name:"Period"`
	// 区域
	Region string `name:"Region"`
	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *RenewInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewInstanceResponse, error) {
	resp := &RenewInstanceResponse{}
	err := client.Request("postgres", "RenewInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type RenewInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单名
	DealName string `json:"DealName"`
}

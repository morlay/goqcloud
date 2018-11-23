package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 升级实例
// https://cloud.tencent.com/document/api/409/18095

type UpgradeDbInstanceRequest struct {
	// 是否自动使用代金券,1是,0否，默认不使用
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 实例ID，形如postgres-lnp6j617
	DBInstanceId string `name:"DBInstanceId"`
	// 升级后的实例内存大小，单位GB
	Memory int64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 升级后的实例磁盘大小，单位GB
	Storage int64 `name:"Storage"`
	// 代金券ID列表，目前仅支持指定一张代金券
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *UpgradeDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeDbInstanceResponse, error) {
	resp := &UpgradeDbInstanceResponse{}
	err := client.Request("postgres", "UpgradeDBInstance", "2017-03-12").Do(req, resp)
	return resp, err
}

type UpgradeDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 交易名字。
	DealName string `json:"DealName"`
}

package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 升级实例
// https://cloud.tencent.com/document/api/238/19948

type UpgradeDbInstanceRequest struct {
	// 是否自动使用代金券，0 - 不使用；1 - 默认使用。取值默认为0
	AutoVoucher *int64 `name:"AutoVoucher,omitempty"`
	// 实例ID，形如mssql-j8kv137v
	InstanceId string `name:"InstanceId"`
	// 实例升级后内存大小，单位GB，其值不能小于当前实例内存大小
	Memory int64 `name:"Memory"`
	// 区域
	Region string `name:"Region"`
	// 实例升级后磁盘大小，单位GB，其值不能小于当前实例磁盘大小
	Storage int64 `name:"Storage"`
	// 代金券ID，目前单个订单只能使用一张代金券
	VoucherIds []*string `name:"VoucherIds,omitempty"`
}

func (req *UpgradeDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeDbInstanceResponse, error) {
	resp := &UpgradeDbInstanceResponse{}
	err := client.Request("sqlserver", "UpgradeDBInstance", "2018-03-28").Do(req, resp)
	return resp, err
}

type UpgradeDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单名称
	DealName string `json:"DealName"`
}

package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例续费标记
// https://cloud.tencent.com/document/api/238/19956

type ModifyDbInstanceRenewFlagRequest struct {
	// 区域
	Region string `name:"Region"`
	// 实例续费状态标记信息
	RenewFlags []*InstanceRenewInfo `name:"RenewFlags"`
}

func (req *ModifyDbInstanceRenewFlagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceRenewFlagResponse, error) {
	resp := &ModifyDbInstanceRenewFlagResponse{}
	err := client.Request("sqlserver", "ModifyDBInstanceRenewFlag", "2018-03-28").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceRenewFlagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 修改成功的个数
	Count int64 `json:"Count"`
}

package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云硬盘续费标识
// https://cloud.tencent.com/document/api/362/15668

type ModifyDisksRenewFlagRequest struct {
	// 一个或多个待操作的云硬盘ID。
	DiskIds []*string `name:"DiskIds"`
	// 区域
	Region string `name:"Region"`
	// 云盘的续费标识。取值范围：NOTIFY_AND_AUTO_RENEW：通知过期且自动续费NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费。
	RenewFlag string `name:"RenewFlag"`
}

func (req *ModifyDisksRenewFlagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDisksRenewFlagResponse, error) {
	resp := &ModifyDisksRenewFlagResponse{}
	err := client.Request("cbs", "ModifyDisksRenewFlag", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDisksRenewFlagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

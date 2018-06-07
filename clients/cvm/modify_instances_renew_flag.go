package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例续费标识
// https://cloud.tencent.com/document/api/213/15752

type ModifyInstancesRenewFlagRequest struct {
	// 一个或多个待操作的实例ID。可通过DescribeInstances API返回值中的InstanceId获取。每次请求允许操作的实例数量上限是100。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
	// 自动续费标识。取值范围：NOTIFY_AND_AUTO_RENEW：通知过期且自动续费NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag string `name:"RenewFlag"`
}

func (req *ModifyInstancesRenewFlagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyInstancesRenewFlagResponse, error) {
	resp := &ModifyInstancesRenewFlagResponse{}
	err := client.Request("cvm", "ModifyInstancesRenewFlag", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyInstancesRenewFlagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

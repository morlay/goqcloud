package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改CDH实例的属性
// https://cloud.tencent.com/document/api/213/16475

type ModifyHostsAttributeRequest struct {
	// 一个或多个待操作的CDH实例ID。
	HostIds []*string `name:"HostIds"`
	// CDH实例显示名称。可任意命名，但不得超过60个字符。
	HostName *string `name:"HostName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 自动续费标识。取值范围：NOTIFY_AND_AUTO_RENEW：通知过期且自动续费NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费若该参数指定为NOTIFY_AND_AUTO_RENEW，在账户余额充足的情况下，实例到期后将按月自动续费。
	RenewFlag *string `name:"RenewFlag,omitempty"`
}

func (req *ModifyHostsAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyHostsAttributeResponse, error) {
	resp := &ModifyHostsAttributeResponse{}
	err := client.Request("cvm", "ModifyHostsAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyHostsAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

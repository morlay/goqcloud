package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库实例的自动续费标记
// https://cloud.tencent.com/document/api/236/19652

type ModifyAutoRenewFlagRequest struct {
	// 自动续费标记，可取值的有：0-不自动续费，1-自动续费。
	AutoRenew int64 `name:"AutoRenew"`
	// 实例的ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyAutoRenewFlagRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAutoRenewFlagResponse, error) {
	resp := &ModifyAutoRenewFlagResponse{}
	err := client.Request("cdb", "ModifyAutoRenewFlag", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyAutoRenewFlagResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

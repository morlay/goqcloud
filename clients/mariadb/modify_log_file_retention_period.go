package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改备份日志保存天数
// https://cloud.tencent.com/document/api/237/16151
type ModifyLogFileRetentionPeriodRequest struct {
	// 保存的天数,不能超过30
	Days int64 `name:"Days"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyLogFileRetentionPeriodRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyLogFileRetentionPeriodResponse, error) {
	resp := &ModifyLogFileRetentionPeriodResponse{}
	err := client.Request("mariadb", "ModifyLogFileRetentionPeriod", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyLogFileRetentionPeriodResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `json:"InstanceId"`
}

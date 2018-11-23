package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置自动备份时间
// https://cloud.tencent.com/document/api/239/20016

type ModifyAutoBackupConfigRequest struct {
	// 自动备份类型： 1 “定时回档”
	AutoBackupType *int64 `name:"AutoBackupType,omitempty"`
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 时间段 00:00-01:00, 01:00-02:00...... 23:00-00:00
	TimePeriod string `name:"TimePeriod"`
	// 日期 Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday
	WeekDays []*string `name:"WeekDays"`
}

func (req *ModifyAutoBackupConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyAutoBackupConfigResponse, error) {
	resp := &ModifyAutoBackupConfigResponse{}
	err := client.Request("redis", "ModifyAutoBackupConfig", "2018-04-12").Do(req, resp)
	return resp, err
}

type ModifyAutoBackupConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 自动备份类型： 1 “定时回档”
	AutoBackupType int64 `json:"AutoBackupType"`
	// 时间段 00:00-01:00, 01:00-02:00...... 23:00-00:00
	TimePeriod string `json:"TimePeriod"`
	// 日期Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
	WeekDays []*string `json:"WeekDays"`
}

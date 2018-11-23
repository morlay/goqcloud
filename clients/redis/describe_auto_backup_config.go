package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取备份配置
// https://cloud.tencent.com/document/api/239/20019

type DescribeAutoBackupConfigRequest struct {
	// 实例ID
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAutoBackupConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAutoBackupConfigResponse, error) {
	resp := &DescribeAutoBackupConfigResponse{}
	err := client.Request("redis", "DescribeAutoBackupConfig", "2018-04-12").Do(req, resp)
	return resp, err
}

type DescribeAutoBackupConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 备份类型。自动备份类型： 1 “定时回档”
	AutoBackupType int64 `json:"AutoBackupType"`
	// 时间段。
	TimePeriod string `json:"TimePeriod"`
	// Monday，Tuesday，Wednesday，Thursday，Friday，Saturday，Sunday。
	WeekDays []*string `json:"WeekDays"`
}

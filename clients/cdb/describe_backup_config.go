package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库备份配置信息
// https://cloud.tencent.com/document/api/236/15837

type DescribeBackupConfigRequest struct {
	// 实例短实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBackupConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupConfigResponse, error) {
	resp := &DescribeBackupConfigResponse{}
	err := client.Request("cdb", "DescribeBackupConfig", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeBackupConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 备份过期时间，单位为天
	BackupExpireDays int64 `json:"BackupExpireDays"`
	// 备份方式，可能的值为：physical - 物理备份，logical - 逻辑备份
	BackupMethod string `json:"BackupMethod"`
	// Binlog过期时间，单位为天
	BinlogExpireDays int64 `json:"BinlogExpireDays"`
	// 备份开始的最晚时间点，单位为时刻。例如，6 - 凌晨6:00
	StartTimeMax int64 `json:"StartTimeMax"`
	// 备份开始的最早时间点，单位为时刻。例如，2 - 凌晨2:00
	StartTimeMin int64 `json:"StartTimeMin"`
}

package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建云数据库备份
// https://cloud.tencent.com/document/api/236/15844
type CreateBackupRequest struct {
	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备。
	BackupMethod string `name:"BackupMethod"`
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateBackupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateBackupResponse, error) {
	resp := &CreateBackupResponse{}
	err := client.Request("cdb", "CreateBackup", "2017-03-20").Do(req, resp)
	return resp, err
}

type CreateBackupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

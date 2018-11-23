package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除云数据库备份
// https://cloud.tencent.com/document/api/236/15841

type DeleteBackupRequest struct {
	// 备份任务ID。该任务ID为创建云数据库备份接口返回的任务ID。
	BackupId int64 `name:"BackupId"`
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteBackupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteBackupResponse, error) {
	resp := &DeleteBackupResponse{}
	err := client.Request("cdb", "DeleteBackup", "2017-03-20").Do(req, resp)
	return resp, err
}

type DeleteBackupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

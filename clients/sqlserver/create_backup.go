package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建备份
// https://cloud.tencent.com/document/api/238/19946

type CreateBackupRequest struct {
	// 需要备份库名的列表(多库备份才填写)
	DBNames []*string `name:"DBNames,omitempty"`
	// 实例ID，形如mssql-i1z41iwd
	InstanceId *string `name:"InstanceId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 备份策略(0-实例备份 1-多库备份)
	Strategy *int64 `name:"Strategy,omitempty"`
}

func (req *CreateBackupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateBackupResponse, error) {
	resp := &CreateBackupResponse{}
	err := client.Request("sqlserver", "CreateBackup", "2018-03-28").Do(req, resp)
	return resp, err
}

type CreateBackupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务ID
	FlowId int64 `json:"FlowId"`
}

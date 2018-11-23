package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 手动备份Redis实例
// https://cloud.tencent.com/document/api/239/20010

type ManualBackupInstanceRequest struct {
	// 待操作的实例ID，可通过 DescribeInstance接口返回值中的 InstanceId 获取。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 备份的备注信息
	Remark *string `name:"Remark,omitempty"`
}

func (req *ManualBackupInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ManualBackupInstanceResponse, error) {
	resp := &ManualBackupInstanceResponse{}
	err := client.Request("redis", "ManualBackupInstance", "2018-04-12").Do(req, resp)
	return resp, err
}

type ManualBackupInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	TaskId int64 `json:"TaskId"`
}

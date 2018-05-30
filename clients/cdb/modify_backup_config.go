package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改数据库备份配置
// https://cloud.tencent.com/document/api/236/15839
type ModifyBackupConfigRequest struct {
	// 目标备份方法，可选的值：logical - 逻辑冷备，physical - 物理冷备；默认备份方法为 逻辑冷备。
	BackupMethod *string `name:"BackupMethod,omitempty"`
	// 备份过期时间，单位为天，最小值为7天，最大值为732天。
	ExpireDays *int64 `name:"ExpireDays,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 备份时间范围，格式为：02:00-06:00，起点和终点时间目前限制为整点，目前可以选择的范围为： 02:00-06:00，06：00-10：00，10:00-14:00，14:00-18:00，18:00-22:00，22:00-02:00。
	StartTime *string `name:"StartTime,omitempty"`
}

func (req *ModifyBackupConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyBackupConfigResponse, error) {
	resp := &ModifyBackupConfigResponse{}
	err := client.Request("cdb", "ModifyBackupConfig", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyBackupConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

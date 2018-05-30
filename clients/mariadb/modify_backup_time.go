package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改备份时间
// https://cloud.tencent.com/document/api/237/16173
type ModifyBackupTimeRequest struct {
	// 每天备份执行的区间的结束时间，格式 mm:ss，形如 23:59
	EndBackupTime string `name:"EndBackupTime"`
	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 每天备份执行的区间的开始时间，格式 mm:ss，形如 22:00
	StartBackupTime string `name:"StartBackupTime"`
}

func (req *ModifyBackupTimeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyBackupTimeResponse, error) {
	resp := &ModifyBackupTimeResponse{}
	err := client.Request("mariadb", "ModifyBackupTime", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyBackupTimeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设置的状态，0 表示成功
	Status int64 `json:"Status"`
}

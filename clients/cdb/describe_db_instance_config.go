package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库实例的配置信息
// https://cloud.tencent.com/document/api/236/17491

type DescribeDbInstanceConfigRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbInstanceConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstanceConfigResponse, error) {
	resp := &DescribeDbInstanceConfigResponse{}
	err := client.Request("cdb", "DescribeDBInstanceConfig", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbInstanceConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// ECDB第二个从库的配置信息，只有ECDB实例才有这个字段。
	BackupConfig BackupConfig `json:"BackupConfig"`
	// 主库部署方式，主实例属性，可能的返回值：0-单可用部署，1-多可用区部署。
	DeployMode int64 `json:"DeployMode"`
	// 主库数据保护方式，主实例属性，可能的返回值：0-异步复制方式，1-半同步复制方式，2-强同步复制方式。
	ProtectMode int64 `json:"ProtectMode"`
	// 从库的配置信息。
	SlaveConfig SlaveConfig `json:"SlaveConfig"`
	// 主库可用区的正式名称，如ap-shanghai-1。
	Zone string `json:"Zone"`
}

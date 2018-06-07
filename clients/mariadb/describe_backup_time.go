package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询备份时间
// https://cloud.tencent.com/document/api/237/16182

type DescribeBackupTimeRequest struct {
	// 实例ID，形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBackupTimeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBackupTimeResponse, error) {
	resp := &DescribeBackupTimeResponse{}
	err := client.Request("mariadb", "DescribeBackupTime", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeBackupTimeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例备份时间配置信息
	Items DBBackupTimeConfig `json:"Items"`
	// 返回的配置数量
	TotalCount int64 `json:"TotalCount"`
}

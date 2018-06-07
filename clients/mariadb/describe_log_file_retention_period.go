package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看备份日志备份天数
// https://cloud.tencent.com/document/api/237/16152

type DescribeLogFileRetentionPeriodRequest struct {
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeLogFileRetentionPeriodRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLogFileRetentionPeriodResponse, error) {
	resp := &DescribeLogFileRetentionPeriodResponse{}
	err := client.Request("mariadb", "DescribeLogFileRetentionPeriod", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeLogFileRetentionPeriodResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 日志备份天数
	Days int64 `json:"Days"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `json:"InstanceId"`
}

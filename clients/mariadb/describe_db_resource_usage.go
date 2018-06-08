package mariadb

import (
	time "time"

	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看实例资源使用情况
// https://cloud.tencent.com/document/api/237/16158

type DescribeDbResourceUsageRequest struct {
	// 结束日期，格式yyyy-mm-dd
	EndTime time.Time `name:"EndTime"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 开始日期，格式yyyy-mm-dd
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbResourceUsageRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbResourceUsageResponse, error) {
	resp := &DescribeDbResourceUsageResponse{}
	err := client.Request("mariadb", "DescribeDBResourceUsage", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbResourceUsageResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// binlog日志磁盘可用空间,单位GB
	BinlogDiskAvailable MonitorData `json:"BinlogDiskAvailable"`
	// CPU利用率
	CpuUsageRate MonitorData `json:"CpuUsageRate"`
	// 磁盘可用空间,单位GB
	DataDiskAvailable MonitorData `json:"DataDiskAvailable"`
	// 内存可用空间,单位GB
	MemAvailable MonitorData `json:"MemAvailable"`
}

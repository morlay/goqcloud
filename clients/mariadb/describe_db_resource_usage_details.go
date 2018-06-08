package mariadb

import (
	time "time"

	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看实例资源使用详情
// https://cloud.tencent.com/document/api/237/16157

type DescribeDbResourceUsageDetailsRequest struct {
	// 结束日期，格式yyyy-mm-dd
	EndTime time.Time `name:"EndTime"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 开始日期，格式yyyy-mm-dd
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbResourceUsageDetailsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbResourceUsageDetailsResponse, error) {
	resp := &DescribeDbResourceUsageDetailsResponse{}
	err := client.Request("mariadb", "DescribeDBResourceUsageDetails", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbResourceUsageDetailsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 主节点资源使用情况监控数据
	Master ResourceUsageMonitorSet `json:"Master"`
	// 备机1资源使用情况监控数据
	Slave1 ResourceUsageMonitorSet `json:"Slave1"`
	// 备机2资源使用情况监控数据
	Slave2 ResourceUsageMonitorSet `json:"Slave2"`
}

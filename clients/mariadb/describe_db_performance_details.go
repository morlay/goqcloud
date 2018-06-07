package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查看实例性能数据详情
// https://cloud.tencent.com/document/api/237/16156

type DescribeDbPerformanceDetailsRequest struct {
	// 结束日期，格式yyyy-mm-dd
	EndTime time.Time `name:"EndTime"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 开始日期，格式yyyy-mm-dd
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbPerformanceDetailsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbPerformanceDetailsResponse, error) {
	resp := &DescribeDbPerformanceDetailsResponse{}
	err := client.Request("mariadb", "DescribeDBPerformanceDetails", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbPerformanceDetailsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 主节点性能监控数据
	Master PerformanceMonitorSet `json:"Master"`
	// 备机1性能监控数据
	Slave1 PerformanceMonitorSet `json:"Slave1"`
	// 备机2性能监控数据，如果实例是一主一从，则没有该字段
	Slave2 PerformanceMonitorSet `json:"Slave2"`
}

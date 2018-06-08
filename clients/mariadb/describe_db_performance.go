package mariadb

import (
	time "time"

	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看实例性能数据
// https://cloud.tencent.com/document/api/237/16160

type DescribeDbPerformanceRequest struct {
	// 结束日期，格式yyyy-mm-dd
	EndTime time.Time `name:"EndTime"`
	// 实例 ID，形如：tdsql-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 开始日期，格式yyyy-mm-dd
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeDbPerformanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbPerformanceResponse, error) {
	resp := &DescribeDbPerformanceResponse{}
	err := client.Request("mariadb", "DescribeDBPerformance", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDbPerformanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 活跃连接数
	ConnActive MonitorData `json:"ConnActive"`
	// 删除操作数DELETE
	DeleteTotal MonitorData `json:"DeleteTotal"`
	// 磁盘每秒IO次数
	DiskIops MonitorData `json:"DiskIops"`
	// 插入操作数INSERT
	InsertTotal MonitorData `json:"InsertTotal"`
	// 是否发生主备切换，1为发生，0否
	IsMasterSwitched MonitorData `json:"IsMasterSwitched"`
	// 慢查询数
	LongQuery MonitorData `json:"LongQuery"`
	// 缓存命中率
	MemHitRate MonitorData `json:"MemHitRate"`
	// 查询操作数SELECT
	SelectTotal MonitorData `json:"SelectTotal"`
	// 主备延迟
	SlaveDelay MonitorData `json:"SlaveDelay"`
	// 更新操作数UPDATE
	UpdateTotal MonitorData `json:"UpdateTotal"`
}

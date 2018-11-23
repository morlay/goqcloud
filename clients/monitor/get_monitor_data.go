package monitor

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 拉取指标监控数据
// https://cloud.tencent.com/document/api/248/30353

type GetMonitorDataRequest struct {
	// 结束时间，默认为当前时间。 EndTime不能小于EtartTime
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 实例对象的维度组合
	Instances []*Instance `name:"Instances"`
	// 指标名称
	MetricName string `name:"MetricName"`
	// 命名空间，每个云产品会有一个命名空间
	Namespace string `name:"Namespace"`
	// 监控统计周期。默认为取值为300，单位为s
	Period *int64 `name:"Period,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 起始时间，如 2018-01-01 00:00:00
	StartTime *time.Time `name:"StartTime,omitempty"`
}

func (req *GetMonitorDataRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetMonitorDataResponse, error) {
	resp := &GetMonitorDataResponse{}
	err := client.Request("monitor", "GetMonitorData", "2018-07-24").Do(req, resp)
	return resp, err
}

type GetMonitorDataResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 数据点数组
	DataPoints []*DataPoint `json:"DataPoints"`
	// 结束时间
	EndTime time.Time `json:"EndTime"`
	// 指标名
	MetricName string `json:"MetricName"`
	// 统计周期
	Period int64 `json:"Period"`
	// 开始时间
	StartTime time.Time `json:"StartTime"`
}

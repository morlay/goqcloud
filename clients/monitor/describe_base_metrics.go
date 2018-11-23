package monitor

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取基础指标详情
// https://cloud.tencent.com/document/api/248/30351

type DescribeBaseMetricsRequest struct {
	// 指标名
	MetricName *string `name:"MetricName,omitempty"`
	// 业务命名空间
	Namespace string `name:"Namespace"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeBaseMetricsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeBaseMetricsResponse, error) {
	resp := &DescribeBaseMetricsResponse{}
	err := client.Request("monitor", "DescribeBaseMetrics", "2018-07-24").Do(req, resp)
	return resp, err
}

type DescribeBaseMetricsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 查询得到的指标描述列表
	MetricSet []*MetricSet `json:"MetricSet"`
}

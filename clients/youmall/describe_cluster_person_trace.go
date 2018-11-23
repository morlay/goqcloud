package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询客户单次到场轨迹（按天聚合）
// https://cloud.tencent.com/document/api/860/30167

type DescribeClusterPersonTraceRequest struct {
	// 查询结束时间
	EndTime time.Time `name:"EndTime"`
	// 卖场编码
	MallId string `name:"MallId"`
	// 客户编码
	PersonId string `name:"PersonId"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartTime time.Time `name:"StartTime"`
}

func (req *DescribeClusterPersonTraceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeClusterPersonTraceResponse, error) {
	resp := &DescribeClusterPersonTraceResponse{}
	err := client.Request("youmall", "DescribeClusterPersonTrace", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeClusterPersonTraceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 卖场用户编码
	MallCode string `json:"MallCode"`
	// 卖场系统编码
	MallId string `json:"MallId"`
	// 客户编码
	PersonId string `json:"PersonId"`
	// 轨迹序列
	TracePointSet []*DailyTracePoint `json:"TracePointSet"`
}

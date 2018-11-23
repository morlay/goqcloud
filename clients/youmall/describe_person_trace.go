package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询客户单次到场轨迹
// https://cloud.tencent.com/document/api/860/30164

type DescribePersonTraceRequest struct {
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

func (req *DescribePersonTraceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePersonTraceResponse, error) {
	resp := &DescribePersonTraceResponse{}
	err := client.Request("youmall", "DescribePersonTrace", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribePersonTraceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 卖场用户编码
	MallCode string `json:"MallCode"`
	// 卖场系统编码
	MallId string `json:"MallId"`
	// 客户编码
	PersonId string `json:"PersonId"`
	// 轨迹列表
	TraceRouteSet []*PersonTraceRoute `json:"TraceRouteSet"`
}

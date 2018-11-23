package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询客户单次到场轨迹明细
// https://cloud.tencent.com/document/api/860/30163

type DescribePersonTraceDetailRequest struct {
	// 卖场编码
	MallId string `name:"MallId"`
	// 客户编码
	PersonId string `name:"PersonId"`
	// 区域
	Region string `name:"Region"`
	// 轨迹编码
	TraceId string `name:"TraceId"`
}

func (req *DescribePersonTraceDetailRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePersonTraceDetailResponse, error) {
	resp := &DescribePersonTraceDetailResponse{}
	err := client.Request("youmall", "DescribePersonTraceDetail", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribePersonTraceDetailResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 轨迹点坐标序列
	CoordinateSet []*PersonCoordinate `json:"CoordinateSet"`
	// 卖场编码
	MallId string `json:"MallId"`
	// 客户编码
	PersonId string `json:"PersonId"`
	// 轨迹编码
	TraceId string `json:"TraceId"`
}

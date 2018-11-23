package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询客户到场信息（按天聚合）
// https://cloud.tencent.com/document/api/860/30168

type DescribeClusterPersonArrivedMallRequest struct {
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

func (req *DescribeClusterPersonArrivedMallRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeClusterPersonArrivedMallResponse, error) {
	resp := &DescribeClusterPersonArrivedMallResponse{}
	err := client.Request("youmall", "DescribeClusterPersonArrivedMall", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeClusterPersonArrivedMallResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 到场信息
	ArrivedMallSet []*ArrivedMallInfo `json:"ArrivedMallSet"`
	// 卖场客户编码
	MallCode string `json:"MallCode"`
	// 卖场系统编码
	MallId string `json:"MallId"`
	// 客户编码
	PersonId string `json:"PersonId"`
}

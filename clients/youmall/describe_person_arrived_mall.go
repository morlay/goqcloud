package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询客户到场信息
// https://cloud.tencent.com/document/api/860/30165

type DescribePersonArrivedMallRequest struct {
	// 查询结束时间
	EndTime string `name:"EndTime"`
	// 卖场编码
	MallId string `name:"MallId"`
	// 客户编码
	PersonId string `name:"PersonId"`
	// 区域
	Region string `name:"Region"`
	// 查询开始时间
	StartTime string `name:"StartTime"`
}

func (req *DescribePersonArrivedMallRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePersonArrivedMallResponse, error) {
	resp := &DescribePersonArrivedMallResponse{}
	err := client.Request("youmall", "DescribePersonArrivedMall", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribePersonArrivedMallResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 到场轨迹
	ArrivedMallSet []*ArrivedMallInfo `json:"ArrivedMallSet"`
	// 卖场用户编码
	MallCode string `json:"MallCode"`
	// 卖场系统编码
	MallId string `json:"MallId"`
	// 客户编码
	PersonId string `json:"PersonId"`
}

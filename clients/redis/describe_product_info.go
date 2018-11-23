package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询产品售卖规格
// https://cloud.tencent.com/document/api/239/30600

type DescribeProductInfoRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeProductInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProductInfoResponse, error) {
	resp := &DescribeProductInfoResponse{}
	err := client.Request("redis", "DescribeProductInfo", "2018-04-12").Do(req, resp)
	return resp, err
}

type DescribeProductInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 地域售卖信息
	RegionSet []*RegionConf `json:"RegionSet"`
}

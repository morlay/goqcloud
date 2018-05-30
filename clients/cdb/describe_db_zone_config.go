package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取云数据库可售卖规格
// https://cloud.tencent.com/document/api/236/17229
type DescribeDbZoneConfigRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbZoneConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbZoneConfigResponse, error) {
	resp := &DescribeDbZoneConfigResponse{}
	err := client.Request("cdb", "DescribeDBZoneConfig", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbZoneConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 可售卖地域配置详情
	Items []*RegionSellConf `json:"Items"`
	// 可售卖地域配置数量
	TotalCount int64 `json:"TotalCount"`
}

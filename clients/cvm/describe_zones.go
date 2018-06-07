package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询可用区列表
// https://cloud.tencent.com/document/api/213/15707

type DescribeZonesRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeZonesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("cvm", "DescribeZones", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 可用区数量
	TotalCount int64 `json:"TotalCount"`
	// 可用区列表信息
	ZoneSet []*ZoneInfo `json:"ZoneSet"`
}

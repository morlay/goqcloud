package postgres

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖可用区
// https://cloud.tencent.com/document/api/409/16769
type DescribeZonesRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeZonesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("postgres", "DescribeZones", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的结果数量。
	TotalCount int64 `json:"TotalCount"`
	// 可用区信息集合。
	ZoneSet []*ZoneInfo `json:"ZoneSet"`
}

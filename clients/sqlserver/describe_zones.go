package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询售卖可用区
// https://cloud.tencent.com/document/api/238/19963

type DescribeZonesRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeZonesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeZonesResponse, error) {
	resp := &DescribeZonesResponse{}
	err := client.Request("sqlserver", "DescribeZones", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeZonesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回多少个可用区信息
	TotalCount int64 `json:"TotalCount"`
	// 可用区数组
	ZoneSet []*ZoneInfo `json:"ZoneSet"`
}

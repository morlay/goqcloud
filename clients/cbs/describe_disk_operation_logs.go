package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云盘操作日志列表
// https://cloud.tencent.com/document/api/362/30162

type DescribeDiskOperationLogsRequest struct {
	// 过滤条件。支持以下条件：disk-id - Array of String - 是否必填：是 - 按云盘ID过滤，每个请求最多可指定10个云盘ID。
	Filters []*Filter `name:"Filters"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDiskOperationLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDiskOperationLogsResponse, error) {
	resp := &DescribeDiskOperationLogsResponse{}
	err := client.Request("cbs", "DescribeDiskOperationLogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDiskOperationLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云盘的操作日志列表。
	DiskOperationLogSet []*DiskOperationLog `json:"DiskOperationLogSet"`
}

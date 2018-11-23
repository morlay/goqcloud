package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 批量查询扫描结果
// https://cloud.tencent.com/document/api/283/17755

type DescribeScanResultsRequest struct {
	// 批量查询一个或者多个app的扫描结果，如果不传表示查询该任务下所提交的所有app
	AppMd5s []*string `name:"AppMd5s,omitempty"`
	// 任务唯一标识
	ItemId string `name:"ItemId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeScanResultsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeScanResultsResponse, error) {
	resp := &DescribeScanResultsResponse{}
	err := client.Request("ms", "DescribeScanResults", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeScanResultsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 批量扫描的app结果集
	ScanSet []*ScanSetInfo `json:"ScanSet"`
	// 批量扫描结果的个数
	TotalCount int64 `json:"TotalCount"`
}

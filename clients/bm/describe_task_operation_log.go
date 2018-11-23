package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 维修任务操作日志获取
// https://cloud.tencent.com/document/api/386/18646

type DescribeTaskOperationLogRequest struct {
	// 排序方式 0:递增(默认) 1:递减
	Order *int64 `name:"Order,omitempty"`
	// 排序字段，目前支持：OperationTime
	OrderField *string `name:"OrderField,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 维修任务ID
	TaskId string `name:"TaskId"`
}

func (req *DescribeTaskOperationLogRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskOperationLogResponse, error) {
	resp := &DescribeTaskOperationLogResponse{}
	err := client.Request("bm", "DescribeTaskOperationLog", "2018-04-23").Do(req, resp)
	return resp, err
}

type DescribeTaskOperationLogResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 操作日志
	TaskOperationLogSet []*TaskOperationLog `json:"TaskOperationLogSet"`
	// 日志条数
	TotalCount int64 `json:"TotalCount"`
}

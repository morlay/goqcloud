package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 维修任务信息获取
// https://cloud.tencent.com/document/api/386/18647

type DescribeTaskInfoRequest struct {
	// 实例别名过滤
	Aliases []*string `name:"Aliases,omitempty"`
	// 时间过滤上限
	EndDate *time.Time `name:"EndDate,omitempty"`
	// 实例ID过滤
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 数据条数
	Limit int64 `name:"Limit"`
	// 开始位置
	Offset int64 `name:"Offset"`
	// 排序方式 0:递增(默认) 1:递减
	Order *int64 `name:"Order,omitempty"`
	// 排序字段，目前支持：CreateTime，AuthTime，EndTime
	OrderField *string `name:"OrderField,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 时间过滤下限
	StartDate *time.Time `name:"StartDate,omitempty"`
	// 任务ID过滤
	TaskIds []*string `name:"TaskIds,omitempty"`
	// 任务状态ID过滤
	TaskStatus []*int64 `name:"TaskStatus,omitempty"`
	// 故障类型ID过滤
	TaskTypeIds []*int64 `name:"TaskTypeIds,omitempty"`
}

func (req *DescribeTaskInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskInfoResponse, error) {
	resp := &DescribeTaskInfoResponse{}
	err := client.Request("bm", "DescribeTaskInfo", "2018-04-23").Do(req, resp)
	return resp, err
}

type DescribeTaskInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务信息列表
	TaskInfoSet []*TaskInfo `json:"TaskInfoSet"`
	// 返回任务总数量
	TotalCount int64 `json:"TotalCount"`
}

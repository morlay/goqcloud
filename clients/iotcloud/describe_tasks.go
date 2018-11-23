package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取任务列表
// https://cloud.tencent.com/document/api/634/19481

type DescribeTasksRequest struct {
	// 分页的大小，数值范围 1-250
	Limit int64 `name:"Limit"`
	// 分页偏移，从0开始
	Offset int64 `name:"Offset"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeTasksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTasksResponse, error) {
	resp := &DescribeTasksResponse{}
	err := client.Request("iotcloud", "DescribeTasks", "2018-06-14").Do(req, resp)
	return resp, err
}

type DescribeTasksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 此页任务对象的数组，按创建时间排序
	Tasks []*TaskInfo `json:"Tasks"`
	// 用户一个月内创建的任务总数
	TotalCount int64 `json:"TotalCount"`
}

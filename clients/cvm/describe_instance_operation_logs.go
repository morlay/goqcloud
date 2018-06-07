package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例操作记录
// https://cloud.tencent.com/document/api/213/15737

type DescribeInstanceOperationLogsRequest struct {
	// 每次请求的Filters的上限为1，Filter.Values的上限为1。Filters.1.Name目前支持“instance-id”，即根据实例 ID 过滤。实例 ID 形如：ins-1w2x3y4z。
	Filters []*Filter `name:"Filters"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstanceOperationLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceOperationLogsResponse, error) {
	resp := &DescribeInstanceOperationLogsResponse{}
	err := client.Request("cvm", "DescribeInstanceOperationLogs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceOperationLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

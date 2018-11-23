package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询异步任务的执行结果
// https://cloud.tencent.com/document/api/236/20410

type DescribeAsyncRequestInfoRequest struct {
	// 异步任务的请求ID。
	AsyncRequestId string `name:"AsyncRequestId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAsyncRequestInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAsyncRequestInfoResponse, error) {
	resp := &DescribeAsyncRequestInfoResponse{}
	err := client.Request("cdb", "DescribeAsyncRequestInfo", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeAsyncRequestInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务执行信息描述。
	Info string `json:"Info"`
	// 任务执行结果。可能的取值：INITIAL - 初始化，RUNNING - 运行中，SUCCESS - 执行成功，FAILED - 执行失败，KILLED - 已终止，REMOVED - 已删除，PAUSED - 终止中。
	Status string `json:"Status"`
}

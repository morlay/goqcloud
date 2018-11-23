package clb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询异步任务状态
// https://cloud.tencent.com/document/api/214/30683

type DescribeTaskStatusRequest struct {
	// 区域
	Region string `name:"Region"`
	// 请求ID，即接口返回的RequestId
	TaskId string `name:"TaskId"`
}

func (req *DescribeTaskStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTaskStatusResponse, error) {
	resp := &DescribeTaskStatusResponse{}
	err := client.Request("clb", "DescribeTaskStatus", "2018-03-17").Do(req, resp)
	return resp, err
}

type DescribeTaskStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务的当前状态。 0：成功，1：失败，2：进行中。
	Status int64 `json:"Status"`
}

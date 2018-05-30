package batch

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取作业的提交信息
// https://cloud.tencent.com/document/api/599/15910
type DescribeJobSubmitInfoRequest struct {
	// 作业ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeJobSubmitInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeJobSubmitInfoResponse, error) {
	resp := &DescribeJobSubmitInfoResponse{}
	err := client.Request("batch", "DescribeJobSubmitInfo", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeJobSubmitInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 依赖信息
	Dependences []*Dependence `json:"Dependences"`
	// 作业描述
	JobDescription string `json:"JobDescription"`
	// 作业ID
	JobId string `json:"JobId"`
	// 作业名称
	JobName string `json:"JobName"`
	// 作业优先级，任务（Task）和任务实例（TaskInstance）会继承作业优先级
	Priority int64 `json:"Priority"`
	// 任务信息
	Tasks []*Task `json:"Tasks"`
}

package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询加固结果
// https://cloud.tencent.com/document/api/283/17750

type DescribeShieldResultRequest struct {
	// 任务唯一标识
	ItemId string `name:"ItemId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeShieldResultRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeShieldResultResponse, error) {
	resp := &DescribeShieldResultResponse{}
	err := client.Request("ms", "DescribeShieldResult", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeShieldResultResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// app加固前的详细信息
	AppDetailInfo AppDetailInfo `json:"AppDetailInfo"`
	// app加固后的详细信息
	ShieldInfo ShieldInfo `json:"ShieldInfo"`
	// 状态描述
	StatusDesc string `json:"StatusDesc"`
	// 状态指引
	StatusRef string `json:"StatusRef"`
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	TaskStatus int64 `json:"TaskStatus"`
}

package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询流程状态
// https://cloud.tencent.com/document/api/237/16177
type DescribeFlowRequest struct {
	// 异步请求接口返回的任务流程号。
	FlowId int64 `name:"FlowId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeFlowRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeFlowResponse, error) {
	resp := &DescribeFlowResponse{}
	err := client.Request("mariadb", "DescribeFlow", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeFlowResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 流程状态，0：成功，1：失败，2：运行中
	Status int64 `json:"Status"`
}

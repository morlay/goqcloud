package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询流程状态
// https://cloud.tencent.com/document/api/238/19967

type DescribeFlowStatusRequest struct {
	// 流程ID
	FlowId int64 `name:"FlowId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeFlowStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeFlowStatusResponse, error) {
	resp := &DescribeFlowStatusResponse{}
	err := client.Request("sqlserver", "DescribeFlowStatus", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeFlowStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 流程状态，0：成功，1：失败，2：运行中
	Status int64 `json:"Status"`
}

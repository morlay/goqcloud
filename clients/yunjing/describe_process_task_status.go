package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取实时拉取进程任务状态
// https://cloud.tencent.com/document/api/296/30333

type DescribeProcessTaskStatusRequest struct {
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。
	Uuid string `name:"Uuid"`
}

func (req *DescribeProcessTaskStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProcessTaskStatusResponse, error) {
	resp := &DescribeProcessTaskStatusResponse{}
	err := client.Request("yunjing", "DescribeProcessTaskStatus", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeProcessTaskStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务状态。COMPLETE：完成（此时可以调用DescribeProcesses接口获取实时进程列表）AGENT_OFFLINE：云镜客户端离线COLLECTING：进程获取中FAILED：进程获取失败
	Status string `json:"Status"`
}

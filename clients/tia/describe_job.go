package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取Job详情
// https://cloud.tencent.com/document/api/851/18306

type DescribeJobRequest struct {
	// 运行任务的集群
	Cluster string `name:"Cluster"`
	// 任务名称
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeJobResponse, error) {
	resp := &DescribeJobResponse{}
	err := client.Request("tia", "DescribeJob", "2018-02-26").Do(req, resp)
	return resp, err
}

type DescribeJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 训练任务信息
	Job Job `json:"Job"`
}

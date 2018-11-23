package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取容器日志信息
// https://cloud.tencent.com/document/api/858/17770

type DescribeContainerLogRequest struct {
	// 容器名称
	ContainerName *string `name:"ContainerName,omitempty"`
	// 容器实例名称
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
	// 日志起始时间
	SinceTime *string `name:"SinceTime,omitempty"`
	// 日志显示尾部行数
	Tail *int64 `name:"Tail,omitempty"`
}

func (req *DescribeContainerLogRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeContainerLogResponse, error) {
	resp := &DescribeContainerLogResponse{}
	err := client.Request("cis", "DescribeContainerLog", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeContainerLogResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 容器日志数组
	ContainerLogList []*ContainerLog `json:"ContainerLogList"`
}

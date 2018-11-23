package cis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询容器实例事件
// https://cloud.tencent.com/document/api/858/17772

type DescribeContainerInstanceEventsRequest struct {
	// 容器实例名称
	InstanceName string `name:"InstanceName"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeContainerInstanceEventsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeContainerInstanceEventsResponse, error) {
	resp := &DescribeContainerInstanceEventsResponse{}
	err := client.Request("cis", "DescribeContainerInstanceEvents", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeContainerInstanceEventsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 容器实例事件列表
	EventList []*Event `json:"EventList"`
}

package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 从伸缩组中删除 CVM 实例
// https://cloud.tencent.com/document/api/377/20431

type RemoveInstancesRequest struct {
	// 伸缩组ID
	AutoScalingGroupId string `name:"AutoScalingGroupId"`
	// CVM实例ID列表
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *RemoveInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RemoveInstancesResponse, error) {
	resp := &RemoveInstancesResponse{}
	err := client.Request("as", "RemoveInstances", "2018-04-19").Do(req, resp)
	return resp, err
}

type RemoveInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

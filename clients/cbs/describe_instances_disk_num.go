package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例关联的云硬盘数量
// https://cloud.tencent.com/document/api/362/16311

type DescribeInstancesDiskNumRequest struct {
	// 云服务器实例ID，通过DescribeInstances接口查询。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstancesDiskNumRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstancesDiskNumResponse, error) {
	resp := &DescribeInstancesDiskNumResponse{}
	err := client.Request("cbs", "DescribeInstancesDiskNum", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstancesDiskNumResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 各个云服务器已挂载和可挂载弹性云盘的数量。
	AttachDetail []*AttachDetail `json:"AttachDetail"`
}

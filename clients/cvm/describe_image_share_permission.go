package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看镜像分享信息
// https://cloud.tencent.com/document/api/213/15712

type DescribeImageSharePermissionRequest struct {
	// 需要共享的镜像Id
	ImageId string `name:"ImageId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeImageSharePermissionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeImageSharePermissionResponse, error) {
	resp := &DescribeImageSharePermissionResponse{}
	err := client.Request("cvm", "DescribeImageSharePermission", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeImageSharePermissionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 镜像共享信息
	SharePermissionSet []*SharePermission `json:"SharePermissionSet"`
}

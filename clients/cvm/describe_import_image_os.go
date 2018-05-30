package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询外部导入镜像支持的OS列表
// https://cloud.tencent.com/document/api/213/15718
type DescribeImportImageOsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeImportImageOsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeImportImageOsResponse, error) {
	resp := &DescribeImportImageOsResponse{}
	err := client.Request("cvm", "DescribeImportImageOs", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeImportImageOsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 支持的导入镜像的操作系统类型
	ImportImageOsListSupported []*string `json:"ImportImageOsListSupported"`
	// 支持的导入镜像的操作系统版本
	ImportImageOsVersionSupported []*string `json:"ImportImageOsVersionSupported"`
}

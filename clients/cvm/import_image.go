package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 外部镜像导入
// https://cloud.tencent.com/document/api/213/15717
type ImportImageRequest struct {
	// 导入镜像的操作系统架构，x86_64 或 i386
	Architecture string `name:"Architecture"`
	// 只检查参数，不执行任务
	DryRun *bool `name:"DryRun,omitempty"`
	// 是否强制导入，参考强制导入镜像
	Force *bool `name:"Force,omitempty"`
	// 镜像描述
	ImageDescription *string `name:"ImageDescription,omitempty"`
	// 镜像名称
	ImageName string `name:"ImageName"`
	// 导入镜像存放的cos地址
	ImageUrl string `name:"ImageUrl"`
	// 导入镜像的操作系统类型，通过DescribeImportImageOs获取
	OsType string `name:"OsType"`
	// 导入镜像的操作系统版本，通过DescribeImportImageOs获取
	OsVersion string `name:"OsVersion"`
	// 区域
	Region string `name:"Region"`
}

func (req *ImportImageRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ImportImageResponse, error) {
	resp := &ImportImageResponse{}
	err := client.Request("cvm", "ImportImage", "2017-03-12").Do(req, resp)
	return resp, err
}

type ImportImageResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

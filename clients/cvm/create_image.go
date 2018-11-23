package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建镜像
// https://cloud.tencent.com/document/api/213/16726

type CreateImageRequest struct {
	// DryRun
	DryRun *bool `name:"DryRun,omitempty"`
	// 软关机失败时是否执行强制关机以制作镜像
	ForcePoweroff *string `name:"ForcePoweroff,omitempty"`
	// 镜像描述
	ImageDescription *string `name:"ImageDescription,omitempty"`
	// 镜像名称
	ImageName string `name:"ImageName"`
	// 需要制作镜像的实例ID
	InstanceId string `name:"InstanceId"`
	// 实例处于运行中时，是否允许关机执行制作镜像任务。
	Reboot *string `name:"Reboot,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 创建Windows镜像时是否启用Sysprep
	Sysprep *string `name:"Sysprep,omitempty"`
}

func (req *CreateImageRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateImageResponse, error) {
	resp := &CreateImageResponse{}
	err := client.Request("cvm", "CreateImage", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateImageResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改镜像属性
// https://cloud.tencent.com/document/api/213/15713

type ModifyImageAttributeRequest struct {
	// 设置新的镜像描述；必须满足下列限制：  不得超过60个字符。
	ImageDescription *string `name:"ImageDescription,omitempty"`
	// 镜像ID，形如img-gvbnzy6f。镜像ID可以通过如下方式获取：通过DescribeImages接口返回的ImageId获取。通过镜像控制台获取。
	ImageId string `name:"ImageId"`
	// 设置新的镜像名称；必须满足下列限制：  不得超过20个字符。  镜像名称不能与已有镜像重复。
	ImageName *string `name:"ImageName,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyImageAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyImageAttributeResponse, error) {
	resp := &ModifyImageAttributeResponse{}
	err := client.Request("cvm", "ModifyImageAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyImageAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

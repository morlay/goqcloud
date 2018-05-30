package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除镜像
// https://cloud.tencent.com/document/api/213/15716
type DeleteImagesRequest struct {
	// 准备删除的镜像Id列表
	ImageIds []*string `name:"ImageIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteImagesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteImagesResponse, error) {
	resp := &DeleteImagesResponse{}
	err := client.Request("cvm", "DeleteImages", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteImagesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

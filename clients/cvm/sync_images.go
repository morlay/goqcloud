package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 同步镜像
// https://cloud.tencent.com/document/api/213/15711

type SyncImagesRequest struct {
	// 目的同步地域列表；必须满足限制：不能为源地域，必须是一个合法的Region。暂不支持部分地域同步。具体地域参数请参考Region。
	DestinationRegions []*string `name:"DestinationRegions"`
	// 镜像ID列表 ，镜像ID可以通过如下方式获取：通过DescribeImages接口返回的ImageId获取。通过镜像控制台获取。镜像ID必须满足限制：镜像ID对应的镜像状态必须为NORMAL。镜像大小小于50GB。镜像状态请参考镜像数据表。
	ImageIds []*string `name:"ImageIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *SyncImagesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SyncImagesResponse, error) {
	resp := &SyncImagesResponse{}
	err := client.Request("cvm", "SyncImages", "2017-03-12").Do(req, resp)
	return resp, err
}

type SyncImagesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

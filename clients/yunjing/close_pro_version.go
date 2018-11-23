package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 关闭专业版
// https://cloud.tencent.com/document/api/296/19847

type CloseProVersionRequest struct {
	// 主机唯一标识Uuid。黑石的InstanceId，CVM的Uuid
	Quuid *string `name:"Quuid,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *CloseProVersionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CloseProVersionResponse, error) {
	resp := &CloseProVersionResponse{}
	err := client.Request("yunjing", "CloseProVersion", "2018-02-28").Do(req, resp)
	return resp, err
}

type CloseProVersionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

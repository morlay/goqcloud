package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 卸载云镜客户端
// https://cloud.tencent.com/document/api/296/19844

type DeleteMachineRequest struct {
	// 区域
	Region string `name:"Region"`
	// 云镜客户端Uuid。
	Uuid string `name:"Uuid"`
}

func (req *DeleteMachineRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteMachineResponse, error) {
	resp := &DeleteMachineResponse{}
	err := client.Request("yunjing", "DeleteMachine", "2018-02-28").Do(req, resp)
	return resp, err
}

type DeleteMachineResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

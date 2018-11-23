package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 添加常用登录地
// https://cloud.tencent.com/document/api/296/19846

type CreateUsualLoginPlacesRequest struct {
	// 登录地域信息数组。
	Places []*Place `name:"Places"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端UUID数组。
	Uuids []*string `name:"Uuids"`
}

func (req *CreateUsualLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateUsualLoginPlacesResponse, error) {
	resp := &CreateUsualLoginPlacesResponse{}
	err := client.Request("yunjing", "CreateUsualLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type CreateUsualLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

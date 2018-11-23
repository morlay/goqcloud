package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除异地登录记录
// https://cloud.tencent.com/document/api/296/19843

type DeleteNonlocalLoginPlacesRequest struct {
	// 异地登录事件Id数组。
	Ids []*int64 `name:"Ids"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteNonlocalLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteNonlocalLoginPlacesResponse, error) {
	resp := &DeleteNonlocalLoginPlacesResponse{}
	err := client.Request("yunjing", "DeleteNonlocalLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type DeleteNonlocalLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

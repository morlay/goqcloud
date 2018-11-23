package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除常用登录地
// https://cloud.tencent.com/document/api/296/19842

type DeleteUsualLoginPlacesRequest struct {
	// 已添加常用登录地城市ID数组
	CityIds []*int64 `name:"CityIds"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端Uuid
	Uuid string `name:"Uuid"`
}

func (req *DeleteUsualLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteUsualLoginPlacesResponse, error) {
	resp := &DeleteUsualLoginPlacesResponse{}
	err := client.Request("yunjing", "DeleteUsualLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type DeleteUsualLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

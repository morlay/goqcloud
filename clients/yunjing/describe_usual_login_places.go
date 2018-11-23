package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询常用登录地
// https://cloud.tencent.com/document/api/296/19841

type DescribeUsualLoginPlacesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 云镜客户端UUID
	Uuid string `name:"Uuid"`
}

func (req *DescribeUsualLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeUsualLoginPlacesResponse, error) {
	resp := &DescribeUsualLoginPlacesResponse{}
	err := client.Request("yunjing", "DescribeUsualLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeUsualLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 常用登录地数组
	UsualLoginPlaces []*UsualPlace `json:"UsualLoginPlaces"`
}

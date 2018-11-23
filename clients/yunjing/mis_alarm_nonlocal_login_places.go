package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 误报异地登录
// https://cloud.tencent.com/document/api/296/19835

type MisAlarmNonlocalLoginPlacesRequest struct {
	// 异地登录事件Id数组。
	Ids []*int64 `name:"Ids"`
	// 区域
	Region string `name:"Region"`
}

func (req *MisAlarmNonlocalLoginPlacesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*MisAlarmNonlocalLoginPlacesResponse, error) {
	resp := &MisAlarmNonlocalLoginPlacesResponse{}
	err := client.Request("yunjing", "MisAlarmNonlocalLoginPlaces", "2018-02-28").Do(req, resp)
	return resp, err
}

type MisAlarmNonlocalLoginPlacesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除暴力破解记录
// https://cloud.tencent.com/document/api/296/19845

type DeleteBruteAttacksRequest struct {
	// 暴力破解事件Id数组。
	Ids []*int64 `name:"Ids"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteBruteAttacksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteBruteAttacksResponse, error) {
	resp := &DeleteBruteAttacksResponse{}
	err := client.Request("yunjing", "DeleteBruteAttacks", "2018-02-28").Do(req, resp)
	return resp, err
}

type DeleteBruteAttacksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 清空Redis实例
// https://cloud.tencent.com/document/api/239/20021

type ClearInstanceRequest struct {
	// 实例id
	InstanceId string `name:"InstanceId"`
	// redis的实例密码
	Password string `name:"Password"`
	// 区域
	Region string `name:"Region"`
}

func (req *ClearInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ClearInstanceResponse, error) {
	resp := &ClearInstanceResponse{}
	err := client.Request("redis", "ClearInstance", "2018-04-12").Do(req, resp)
	return resp, err
}

type ClearInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务Id
	TaskId int64 `json:"TaskId"`
}

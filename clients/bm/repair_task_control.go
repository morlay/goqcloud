package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 维修任务管理
// https://cloud.tencent.com/document/api/386/18645

type RepairTaskControlRequest struct {
	// 操作
	Operate string `name:"Operate"`
	// 区域
	Region string `name:"Region"`
	// 维修任务ID
	TaskId string `name:"TaskId"`
}

func (req *RepairTaskControlRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RepairTaskControlResponse, error) {
	resp := &RepairTaskControlResponse{}
	err := client.Request("bm", "RepairTaskControl", "2018-04-23").Do(req, resp)
	return resp, err
}

type RepairTaskControlResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 出参TaskId是黑石异步任务ID，不同于入参TaskId字段。此字段可作为DescriptionOperationResult查询异步任务状态接口的入参，查询异步任务执行结果。
	TaskId int64 `json:"TaskId"`
}

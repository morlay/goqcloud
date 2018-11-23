package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 维修任务配置获取
// https://cloud.tencent.com/document/api/386/18648

type DescribeRepairTaskConstantRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeRepairTaskConstantRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeRepairTaskConstantResponse, error) {
	resp := &DescribeRepairTaskConstantResponse{}
	err := client.Request("bm", "DescribeRepairTaskConstant", "2018-04-23").Do(req, resp)
	return resp, err
}

type DescribeRepairTaskConstantResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 故障类型ID与对应中文名列表
	TaskTypeSet []*TaskType `json:"TaskTypeSet"`
}

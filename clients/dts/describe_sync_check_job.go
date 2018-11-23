package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询灾备同步任务校验结果
// https://cloud.tencent.com/document/api/571/18571

type DescribeSyncCheckJobRequest struct {
	// 要查询的灾备同步任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeSyncCheckJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSyncCheckJobResponse, error) {
	resp := &DescribeSyncCheckJobResponse{}
	err := client.Request("dts", "DescribeSyncCheckJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type DescribeSyncCheckJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 校验标志
	CheckFlag int64 `json:"CheckFlag"`
	// 任务校验结果代码
	ErrorCode int64 `json:"ErrorCode"`
	// 提示信息
	ErrorMessage string `json:"ErrorMessage"`
	// 任务校验状态
	Status string `json:"Status"`
	// 任务执行步骤描述
	StepInfo []*SyncCheckStepInfo `json:"StepInfo"`
}

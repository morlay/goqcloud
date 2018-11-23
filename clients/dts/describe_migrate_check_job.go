package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取迁移校验结果
// https://cloud.tencent.com/document/api/571/18139

type DescribeMigrateCheckJobRequest struct {
	// 数据迁移任务ID
	JobId string `name:"JobId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeMigrateCheckJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMigrateCheckJobResponse, error) {
	resp := &DescribeMigrateCheckJobResponse{}
	err := client.Request("dts", "DescribeMigrateCheckJob", "2018-03-30").Do(req, resp)
	return resp, err
}

type DescribeMigrateCheckJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 校验是否通过,0-未通过，1-校验通过, 3-未校验
	CheckFlag int64 `json:"CheckFlag"`
	// 任务的错误码
	ErrorCode int64 `json:"ErrorCode"`
	// 任务的错误信息
	ErrorMessage string `json:"ErrorMessage"`
	// Check任务总进度,如："30"表示30%
	Progress string `json:"Progress"`
	// 校验任务状态：unavailable(当前不可用), starting(开始中)，running(校验中)，finished(校验完成)
	Status string `json:"Status"`
}

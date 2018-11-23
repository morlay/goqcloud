package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 终止录制任务
// https://cloud.tencent.com/document/api/267/30146

type StopLiveRecordRequest struct {
	// 区域
	Region string `name:"Region"`
	// 流名称。
	StreamName string `name:"StreamName"`
	// 任务ID，全局唯一标识录制任务。
	TaskId int64 `name:"TaskId"`
}

func (req *StopLiveRecordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*StopLiveRecordResponse, error) {
	resp := &StopLiveRecordResponse{}
	err := client.Request("live", "StopLiveRecord", "2018-08-01").Do(req, resp)
	return resp, err
}

type StopLiveRecordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

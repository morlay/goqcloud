package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除录制任务
// https://cloud.tencent.com/document/api/267/30147

type DeleteLiveRecordRequest struct {
	// 区域
	Region string `name:"Region"`
	// 流名称。
	StreamName string `name:"StreamName"`
	// 任务ID，全局唯一标识录制任务。
	TaskId int64 `name:"TaskId"`
}

func (req *DeleteLiveRecordRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteLiveRecordResponse, error) {
	resp := &DeleteLiveRecordResponse{}
	err := client.Request("live", "DeleteLiveRecord", "2018-08-01").Do(req, resp)
	return resp, err
}

type DeleteLiveRecordResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

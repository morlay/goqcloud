package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除定时任务
// https://cloud.tencent.com/document/api/377/20451

type DeleteScheduledActionRequest struct {
	// 区域
	Region string `name:"Region"`
	// 待删除的定时任务ID。
	ScheduledActionId string `name:"ScheduledActionId"`
}

func (req *DeleteScheduledActionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteScheduledActionResponse, error) {
	resp := &DeleteScheduledActionResponse{}
	err := client.Request("as", "DeleteScheduledAction", "2018-04-19").Do(req, resp)
	return resp, err
}

type DeleteScheduledActionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

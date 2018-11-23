package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改同步模式
// https://cloud.tencent.com/document/api/557/19981

type ModifyDbSyncModeRequest struct {
	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode int64 `name:"SyncMode"`
}

func (req *ModifyDbSyncModeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbSyncModeResponse, error) {
	resp := &ModifyDbSyncModeResponse{}
	err := client.Request("dcdb", "ModifyDBSyncMode", "2018-04-11").Do(req, resp)
	return resp, err
}

type ModifyDbSyncModeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowId int64 `json:"FlowId"`
}

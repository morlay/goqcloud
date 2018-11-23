package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询同步模式
// https://cloud.tencent.com/document/api/557/19991

type DescribeDbSyncModeRequest struct {
	// 待修改同步模式的实例ID。形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeDbSyncModeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbSyncModeResponse, error) {
	resp := &DescribeDbSyncModeResponse{}
	err := client.Request("dcdb", "DescribeDBSyncMode", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDbSyncModeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 是否有修改流程在执行中：1 是， 0 否。
	IsModifying int64 `json:"IsModifying"`
	// 同步模式：0 异步，1 强同步， 2 强同步可退化
	SyncMode int64 `json:"SyncMode"`
}

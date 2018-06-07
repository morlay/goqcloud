package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改快照信息
// https://cloud.tencent.com/document/api/362/15650

type ModifySnapshotAttributeRequest struct {
	// 快照的保留时间，FALSE表示非永久保留，TRUE表示永久保留。仅支持将非永久快照修改为永久快照。
	IsPermanent *bool `name:"IsPermanent,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 快照ID, 可通过DescribeSnapshots查询。
	SnapshotId string `name:"SnapshotId"`
	// 新的快照名称。最长为60个字符。
	SnapshotName *string `name:"SnapshotName,omitempty"`
}

func (req *ModifySnapshotAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifySnapshotAttributeResponse, error) {
	resp := &ModifySnapshotAttributeResponse{}
	err := client.Request("cbs", "ModifySnapshotAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifySnapshotAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

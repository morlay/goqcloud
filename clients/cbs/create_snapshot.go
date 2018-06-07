package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建快照
// https://cloud.tencent.com/document/api/362/15648

type CreateSnapshotRequest struct {
	// 需要创建快照的云硬盘ID，可通过DescribeDisks接口查询。
	DiskId string `name:"DiskId"`
	// 区域
	Region string `name:"Region"`
	// 快照名称，不传则新快照名称默认为“未命名”。
	SnapshotName *string `name:"SnapshotName,omitempty"`
}

func (req *CreateSnapshotRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSnapshotResponse, error) {
	resp := &CreateSnapshotResponse{}
	err := client.Request("cbs", "CreateSnapshot", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateSnapshotResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 新创建的快照ID。
	SnapshotId string `json:"SnapshotId"`
}

package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 回滚快照
// https://cloud.tencent.com/document/api/362/15643
type ApplySnapshotRequest struct {
	// 快照原云硬盘ID，可通过DescribeDisks接口查询。
	DiskId string `name:"DiskId"`
	// 区域
	Region string `name:"Region"`
	// 快照ID, 可通过DescribeSnapshots查询。
	SnapshotId string `name:"SnapshotId"`
}

func (req *ApplySnapshotRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ApplySnapshotResponse, error) {
	resp := &ApplySnapshotResponse{}
	err := client.Request("cbs", "ApplySnapshot", "2017-03-12").Do(req, resp)
	return resp, err
}

type ApplySnapshotResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

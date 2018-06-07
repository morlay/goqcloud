package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除快照
// https://cloud.tencent.com/document/api/362/15645

type DeleteSnapshotsRequest struct {
	// 区域
	Region string `name:"Region"`
	// 要删除的快照ID列表，可通过DescribeSnapshots查询。
	SnapshotIds []*string `name:"SnapshotIds"`
}

func (req *DeleteSnapshotsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSnapshotsResponse, error) {
	resp := &DeleteSnapshotsResponse{}
	err := client.Request("cbs", "DeleteSnapshots", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteSnapshotsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

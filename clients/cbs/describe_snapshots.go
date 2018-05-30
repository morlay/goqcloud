package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询快照列表
// https://cloud.tencent.com/document/api/362/15647
type DescribeSnapshotsRequest struct {
	// 过滤条件。参数不支持同时指定SnapshotIds和Filters。snapshot-id - Array of String - 是否必填：否 -（过滤条件）按照快照的ID过滤。快照ID形如：snap-11112222。snapshot-name - Array of String - 是否必填：否 -（过滤条件）按照快照名称过滤。snapshot-state - Array of String - 是否必填：否 -（过滤条件）按照快照状态过滤。 (NORMAL：正常 | CREATING：创建中 | ROLLBACKING：回滚中。)disk-usage - Array of String - 是否必填：否 -（过滤条件）按创建快照的云盘类型过滤。 (SYSTEM_DISK：代表系统盘 | DATA_DISK：代表数据盘。)project-id  - Array of String - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。disk-id  - Array of String - 是否必填：否 -（过滤条件）按照创建快照的云硬盘ID过滤。zone - Array of String - 是否必填：否 -（过滤条件）按照可用区过滤。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考API简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 输出云盘列表的排列顺序。取值范围：ASC：升序排列DESC：降序排列。
	Order *string `name:"Order,omitempty"`
	// 快照列表排序的依据字段。取值范围：CREATE_TIME：依据快照的创建时间排序默认按创建时间排序。
	OrderField *string `name:"OrderField,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 要查询快照的ID列表。参数不支持同时指定SnapshotIds和Filters。
	SnapshotIds []*string `name:"SnapshotIds,omitempty"`
}

func (req *DescribeSnapshotsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSnapshotsResponse, error) {
	resp := &DescribeSnapshotsResponse{}
	err := client.Request("cbs", "DescribeSnapshots", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSnapshotsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 快照的详情列表。
	SnapshotSet []*Snapshot `json:"SnapshotSet"`
	// 快照的数量。
	TotalCount int64 `json:"TotalCount"`
}

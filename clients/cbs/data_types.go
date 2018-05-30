package cbs

import (
	time "time"
)

// 云盘配置。
type DiskConfig struct {
	Available      bool   `json:"Available"`
	DeviceClass    string `json:"DeviceClass"`
	DiskChargeType string `json:"DiskChargeType"`
	DiskType       string `json:"DiskType"`
	DiskUsage      string `json:"DiskUsage"`
	InstanceFamily string `json:"InstanceFamily"`
	MaxDiskSize    int64  `json:"MaxDiskSize"`
	MinDiskSize    int64  `json:"MinDiskSize"`
	Zone           string `json:"Zone"`
}

// 描述键值对过滤器，用于条件过滤查询。
type Filter struct {
	// 过滤键的名称。
	Name string `json:"Name"`
	// 一个或者多个过滤值。
	Values []*string `json:"Values"`
}

// 描述了实例的抽象位置，包括其所在的可用区，所属的项目
type Placement struct {
	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的 projectId 字段来获取。不填为默认项目。
	ProjectId *int64 `json:"ProjectId,omitempty"`
	// 实例所属的可用区ID。该参数也可以通过调用  DescribeZones 的返回值中的Zone字段来获取。
	Zone string `json:"Zone"`
}

// 描述了云盘的价格
type Price struct {
	DiscountPrice float64 `json:"DiscountPrice"`
	OriginalPrice float64 `json:"OriginalPrice"`
}

// 描述了快照的详细信息
type Snapshot struct {
	CopyFromRemote   bool      `json:"CopyFromRemote"`
	CopyingToRegions []*string `json:"CopyingToRegions"`
	CreateTime       time.Time `json:"CreateTime"`
	DeadlineTime     time.Time `json:"DeadlineTime"`
	DiskId           string    `json:"DiskId"`
	DiskSize         int64     `json:"DiskSize"`
	DiskUsage        string    `json:"DiskUsage"`
	Encrypt          bool      `json:"Encrypt"`
	IsPermanent      bool      `json:"IsPermanent"`
	Percent          int64     `json:"Percent"`
	Placement        Placement `json:"Placement"`
	SnapshotId       string    `json:"SnapshotId"`
	SnapshotName     string    `json:"SnapshotName"`
	SnapshotState    string    `json:"SnapshotState"`
}

// 标签。
type Tag struct {
	// 标签健。
	Key string `json:"Key"`
	// 标签值。
	Value string `json:"Value"`
}

// 描述了云硬盘的详细信息
type Disk struct {
	Attached              bool      `json:"Attached"`
	AutoRenewFlagError    bool      `json:"AutoRenewFlagError"`
	AutoSnapshotPolicyIds []*string `json:"AutoSnapshotPolicyIds"`
	CreateTime            time.Time `json:"CreateTime"`
	DeadlineError         bool      `json:"DeadlineError"`
	DeadlineTime          time.Time `json:"DeadlineTime"`
	DiskChargeType        string    `json:"DiskChargeType"`
	DiskId                string    `json:"DiskId"`
	DiskName              string    `json:"DiskName"`
	DiskSize              int64     `json:"DiskSize"`
	DiskState             string    `json:"DiskState"`
	DiskType              string    `json:"DiskType"`
	DiskUsage             string    `json:"DiskUsage"`
	Encrypt               bool      `json:"Encrypt"`
	InstanceId            string    `json:"InstanceId"`
	IsReturnable          bool      `json:"IsReturnable"`
	Placement             Placement `json:"Placement"`
	Portable              bool      `json:"Portable"`
	RenewFlag             string    `json:"RenewFlag"`
	ReturnFailCode        int64     `json:"ReturnFailCode"`
	RollbackPercent       int64     `json:"RollbackPercent"`
	Rollbacking           bool      `json:"Rollbacking"`
	SnapshotAbility       bool      `json:"SnapshotAbility"`
	Tags                  []*Tag    `json:"Tags"`
}

// 描述了实例的计费模式
type DiskChargePrepaid struct {
	// 需要将云盘的到期时间与挂载的子机对齐时，可传入该参数。该参数表示子机当前的到期时间，此时Period如果传入，则表示子机需要续费的时长，云盘会自动按对齐到子机续费后的到期时间续费。
	CurInstanceDeadline *time.Time `json:"CurInstanceDeadline,omitempty"`
	// 购买云盘的时长，默认单位为月，此时，取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period int64 `json:"Period"`
	// 自动续费标识。取值范围：NOTIFY_AND_AUTO_RENEW：通知过期且自动续费NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费DISABLE_NOTIFY_AND_MANUAL_RENEW：不通知过期不自动续费默认取值：NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。
	RenewFlag *string `json:"RenewFlag,omitempty"`
}

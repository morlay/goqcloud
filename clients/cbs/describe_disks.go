package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云硬盘列表
// https://cloud.tencent.com/document/api/362/16315

type DescribeDisksRequest struct {
	// 按照一个或者多个云硬盘ID查询。云硬盘ID形如：disk-11112222，此参数的具体格式可参考API简介的ids.N一节）。参数不支持同时指定DiskIds和Filters。
	DiskIds []*string `name:"DiskIds,omitempty"`
	// 过滤条件。参数不支持同时指定DiskIds和Filters。disk-usage - Array of String - 是否必填：否 -（过滤条件）按云盘类型过滤。 (SYSTEM_DISK：表示系统盘 | DATA_DISK：表示数据盘)disk-charge-type - Array of String - 是否必填：否 -（过滤条件）按照云硬盘计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费。)portable - Array of String - 是否必填：否 -（过滤条件）按是否为弹性云盘过滤。 (TRUE：表示弹性云盘 | FALSE：表示非弹性云盘。)project-id - Array of Integer - 是否必填：否 -（过滤条件）按云硬盘所属项目ID过滤。disk-id - Array of String - 是否必填：否 -（过滤条件）按照云硬盘ID过滤。云盘ID形如：disk-11112222。disk-name - Array of String - 是否必填：否 -（过滤条件）按照云盘名称过滤。disk-type - Array of String - 是否必填：否 -（过滤条件）按照云盘介质类型过滤。(CLOUD_BASIC：表示普通云硬盘 | CLOUD_PREMIUM：表示高性能云硬盘。| CLOUD_SSD：SSD表示SSD云硬盘。)disk-state - Array of String - 是否必填：否 -（过滤条件）按照云盘状态过滤。(UNATTACHED：未挂载 | ATTACHING：挂载中 | ATTACHED：已挂载 | DETACHING：解挂中 | EXPANDING：扩容中 | ROLLBACKING：回滚中 | TORECYCLE：待回收。)instance-id - Array of String - 是否必填：否 -（过滤条件）按照云盘挂载的云主机实例ID过滤。可根据此参数查询挂载在指定云主机下的云硬盘。zone - Array of String - 是否必填：否 -（过滤条件）按照可用区过滤。instance-ip-address - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载云主机的内网或外网IP过滤。instance-name - Array of String - 是否必填：否 -（过滤条件）按云盘所挂载的实例名称过滤。tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键进行过滤。tag-value - Array of String - 是否必填：否 -（过滤条件）照标签值进行过滤。tag:tag-key - Array of String - 是否必填：否 -（过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考API简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 输出云盘列表的排列顺序。取值范围：ASC：升序排列DESC：降序排列。
	Order *string `name:"Order,omitempty"`
	// 云盘列表排序的依据字段。取值范围：CREATE_TIME：依据云盘的创建时间排序DEADLINE：依据云盘的到期时间排序默认按云盘创建时间排序。
	OrderField *string `name:"OrderField,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云盘详情中是否需要返回云盘绑定的定期快照策略ID，TRUE表示需要返回，FALSE表示不返回。
	ReturnBindAutoSnapshotPolicy *bool `name:"ReturnBindAutoSnapshotPolicy,omitempty"`
}

func (req *DescribeDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDisksResponse, error) {
	resp := &DescribeDisksResponse{}
	err := client.Request("cbs", "DescribeDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 云硬盘的详细信息列表。
	DiskSet []*Disk `json:"DiskSet"`
	// 符合条件的云硬盘数量。
	TotalCount int64 `json:"TotalCount"`
}

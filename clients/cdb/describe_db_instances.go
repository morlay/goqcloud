package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例列表
// https://cloud.tencent.com/document/api/236/15872
type DescribeDbInstancesRequest struct {
	// 是否锁定标记
	CdbErrors []*int64 `name:"CdbErrors,omitempty"`
	// 实例数据库引擎版本，可能取值：5.1、5.5、5.6和5.7
	EngineVersions []*string `name:"EngineVersions,omitempty"`
	// 独享集群ID
	ExClusterId *string `name:"ExClusterId,omitempty"`
	// 初始化标记，可取值：0-未初始化，1-初始化
	InitFlag *int64 `name:"InitFlag,omitempty"`
	// 实例ID
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 实例名称
	InstanceNames []*string `name:"InstanceNames,omitempty"`
	// 实例类型，可取值：1-主实例，2-灾备实例，3-只读实例
	InstanceTypes []*int64 `name:"InstanceTypes,omitempty"`
	// 单次请求返回的数量，默认值为20，最大值为100
	Limit *int64 `name:"Limit,omitempty"`
	// 记录偏移量，默认值为0
	Offset *int64 `name:"Offset,omitempty"`
	// 排序的字段，目前支持："InstanceId", "InstanceName", "CreateTime", "DeadlineTime"
	OrderBy *string `name:"OrderBy,omitempty"`
	// 排序方式，目前支持："ASC"或者"DESC"
	OrderDirection *string `name:"OrderDirection,omitempty"`
	// 付费类型，可取值：0-包年包月，1-小时计费
	PayTypes []*int64 `name:"PayTypes,omitempty"`
	// 项目ID，可使用查询项目列表接口查询项目ID
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 安全组ID
	SecurityGroupId *string `name:"SecurityGroupId,omitempty"`
	// 实例状态，可取值：0-创建中，1-运行中，4-删除中，5-隔离中
	Status []*int64 `name:"Status,omitempty"`
	// 子网ID
	SubnetIds []*int64 `name:"SubnetIds,omitempty"`
	// 实例任务状态，可能取值：0-没有任务1-升级中2-数据导入中3-开放Slave中4-外网访问开通中5-批量操作执行中6-回档中7-外网访问关闭中8-密码修改中9-实例名修改中10-重启中12-自建迁移中13-删除库表中14-灾备实例创建同步中
	TaskStatus []*int64 `name:"TaskStatus,omitempty"`
	// 实例的内网IP地址
	Vips []*string `name:"Vips,omitempty"`
	// 私有网络的ID
	VpcIds []*int64 `name:"VpcIds,omitempty"`
	// 是否包含灾备实例
	WithDr *int64 `name:"WithDr,omitempty"`
	// 是否包含独享集群信息
	WithExCluster *int64 `name:"WithExCluster,omitempty"`
	// 是否包含主实例
	WithMaster *int64 `name:"WithMaster,omitempty"`
	// 是否包含只读实例
	WithRo *int64 `name:"WithRo,omitempty"`
	// 是否包含安全组信息
	WithSecurityGroup *int64 `name:"WithSecurityGroup,omitempty"`
	// 可用区的ID
	ZoneIds []*int64 `name:"ZoneIds,omitempty"`
}

func (req *DescribeDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstancesResponse, error) {
	resp := &DescribeDbInstancesResponse{}
	err := client.Request("cdb", "DescribeDBInstances", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息
	Items []*InstanceInfo `json:"Items"`
	// 符合查询条件的实例总数
	TotalCount int64 `json:"TotalCount"`
}

package sqlserver

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例列表
// https://cloud.tencent.com/document/api/238/19969

type DescribeDbInstancesRequest struct {
	// 一个或者多个实例ID。实例ID，格式如：mssql-si2823jyl
	InstanceIdSet []*string `name:"InstanceIdSet,omitempty"`
	// 返回数量，默认为50
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为 0
	Offset *int64 `name:"Offset,omitempty"`
	// 项目ID
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例状态。取值范围：1：申请中2：运行中3：受限运行中 (主备切换中)4：已隔离5：回收中6：已回收7：任务执行中 (实例做备份、回档等操作)8：已下线9：实例扩容中10：实例迁移中11：只读12：重启中
	Status *int64 `name:"Status,omitempty"`
}

func (req *DescribeDbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDbInstancesResponse, error) {
	resp := &DescribeDbInstancesResponse{}
	err := client.Request("sqlserver", "DescribeDBInstances", "2018-03-28").Do(req, resp)
	return resp, err
}

type DescribeDbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例列表
	DBInstances []*DBInstance `json:"DBInstances"`
	// 符合条件的实例总数。分页返回的话，这个值指的是所有符合条件的实例的个数，而非当前根据Limit和Offset值返回的实例个数
	TotalCount int64 `json:"TotalCount"`
}

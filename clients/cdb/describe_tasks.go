package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询云数据库实例任务列表
// https://cloud.tencent.com/document/api/236/15848
type DescribeTasksRequest struct {
	// 异步任务请求ID，执行 CDB 相关操作返回的 AsyncRequestId
	AsyncRequestId *string `name:"AsyncRequestId,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值
	InstanceId *string `name:"InstanceId,omitempty"`
	// 单次请求返回的数量，默认值为20，最大值为100
	Limit *int64 `name:"Limit,omitempty"`
	// 记录偏移量，默认值为0
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 第一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01
	StartTimeBegin *string `name:"StartTimeBegin,omitempty"`
	// 最后一个任务的开始时间，用于范围查询，时间格式如：2017-12-31 10:40:01
	StartTimeEnd *string `name:"StartTimeEnd,omitempty"`
	// 任务状态，不传值则查询所有任务状态，可能的值：-1-未定义；0-初始化; 1-运行中；2-执行成功；3-执行失败；4-已终止；5-已删除；6-已暂停；
	TaskStatus []*int64 `name:"TaskStatus,omitempty"`
	// 任务类型，不传值则查询所有任务类型，可能的值：1-数据库回档；2-SQL操作；3-数据导入；5-参数设置；6-初始化；7-重启；8-开启GTID；9-只读实例升级；10-数据库批量回档；11-主实例升级；12-删除库表；13-切换为主实例；
	TaskTypes []*int64 `name:"TaskTypes,omitempty"`
}

func (req *DescribeTasksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTasksResponse, error) {
	resp := &DescribeTasksResponse{}
	err := client.Request("cdb", "DescribeTasks", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeTasksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的实例任务信息
	Items []*string `json:"Items"`
	// 符合查询条件的实例总数
	TotalCount int64 `json:"TotalCount"`
}

package redis

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询Redis实例备份列表
// https://cloud.tencent.com/document/api/239/20011

type DescribeInstanceBackupsRequest struct {
	// 开始时间，格式如：2017-02-08 16:46:34。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	BeginTime *string `name:"BeginTime,omitempty"`
	// 结束时间，格式如：2017-02-08 19:09:26。查询实例在 [beginTime, endTime] 时间段内开始备份的备份列表。
	EndTime *string `name:"EndTime,omitempty"`
	// 待操作的实例ID，可通过 DescribeInstance 接口返回值中的 InstanceId 获取。
	InstanceId string `name:"InstanceId"`
	// 实例列表大小
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，取Limit整数倍
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 1：备份在流程中，2：备份正常，3：备份转RDB文件处理中，4：已完成RDB转换，-1：备份已过期，-2：备份已删除。
	Status []*string `name:"Status,omitempty"`
}

func (req *DescribeInstanceBackupsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstanceBackupsResponse, error) {
	resp := &DescribeInstanceBackupsResponse{}
	err := client.Request("redis", "DescribeInstanceBackups", "2018-04-12").Do(req, resp)
	return resp, err
}

type DescribeInstanceBackupsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例的备份数组
	BackupSet []*RedisBackupSet `json:"BackupSet"`
	// 备份总数
	TotalCount int64 `json:"TotalCount"`
}

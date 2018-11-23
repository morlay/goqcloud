package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 查询物理机信息
// https://cloud.tencent.com/document/api/386/30257

type DescribeDevicesRequest struct {
	// 设备名称
	Alias *string `name:"Alias,omitempty"`
	// 自动续费标志 0:不自动续费，1:自动续费
	AutoRenewFlag *int64 `name:"AutoRenewFlag,omitempty"`
	// 设备到期时间查询的结束时间
	DeadlineEndTime *time.Time `name:"DeadlineEndTime,omitempty"`
	// 设备到期时间查询的起始时间
	DeadlineStartTime *time.Time `name:"DeadlineStartTime,omitempty"`
	// 机型ID，通过接口查询设备型号(DescribeDeviceClass)查询
	DeviceClassCode *string `name:"DeviceClassCode,omitempty"`
	// 设备类型，取值有: compute(计算型), standard(标准型), storage(存储型) 等
	DeviceType *string `name:"DeviceType,omitempty"`
	// 设备ID数组
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 竞价实例机器的过滤。如果未指定此参数，则不做过滤。0: 查询非竞价实例的机器; 1: 查询竞价实例的机器。
	IsLuckyDevice *int64 `name:"IsLuckyDevice,omitempty"`
	// 内网IP数组
	LanIps []*string `name:"LanIps,omitempty"`
	// 返回数量
	Limit int64 `name:"Limit"`
	// 偏移量
	Offset int64 `name:"Offset"`
	// 排序方式，取值：0:增序(默认)，1:降序
	Order *int64 `name:"Order,omitempty"`
	// 排序字段
	OrderField *string `name:"OrderField,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 子网唯一ID
	SubnetId *string `name:"SubnetId,omitempty"`
	// 标签列表
	Tags []*Tag `name:"Tags,omitempty"`
	// 模糊IP查询
	VagueIp *string `name:"VagueIp,omitempty"`
	// 私有网络唯一ID
	VpcId *string `name:"VpcId,omitempty"`
	// 外网IP数组
	WanIps []*string `name:"WanIps,omitempty"`
}

func (req *DescribeDevicesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDevicesResponse, error) {
	resp := &DescribeDevicesResponse{}
	err := client.Request("bm", "DescribeDevices", "2018-04-23").Do(req, resp)
	return resp, err
}

type DescribeDevicesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 物理机信息列表
	DeviceInfoSet []*DeviceInfo `json:"DeviceInfoSet"`
	// 返回数量
	TotalCount int64 `json:"TotalCount"`
}

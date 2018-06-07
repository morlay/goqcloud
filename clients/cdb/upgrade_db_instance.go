package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 升级云数据库实例
// https://cloud.tencent.com/document/api/236/15876

type UpgradeDbInstanceRequest struct {
	// 备库2的可用区ID，默认为0，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义
	BackupZone *string `name:"BackupZone,omitempty"`
	// 部署模式，默认为0，支持值包括：0-单可用区部署，1-多可用区部署，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义
	DeployMode *int64 `name:"DeployMode,omitempty"`
	// 主实例数据库引擎版本，支持值包括：5.5、5.6和5.7
	EngineVersion *string `name:"EngineVersion,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值
	InstanceId string `name:"InstanceId"`
	// 实例类型，默认为 master，支持值包括：master-表示主实例，dr-表示灾备实例，ro-表示只读实例
	InstanceRole *string `name:"InstanceRole,omitempty"`
	// 升级后的内存大小，单位：MB，为保证传入 Memory 值有效，请使用查询可创建规格（支持可用区、配置自定义）接口获取可升级的内存规格
	Memory int64 `name:"Memory"`
	// 数据复制方式，支持值包括：0-异步复制，1-半同步复制，2-强同步复制，升级主实例时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义
	ProtectMode *int64 `name:"ProtectMode,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 备库1的可用区信息，默认为实例的Zone，升级主实例为多可用区部署时可指定该参数，升级只读实例或者灾备实例时指定该参数无意义。可通过查询云数据库可售卖规格查询支持的可用区
	SlaveZone *string `name:"SlaveZone,omitempty"`
	// 升级后的硬盘大小，单位：GB，为保证传入 Volume 值有效，请使用查询可创建规格（支持可用区、配置自定义）接口获取可升级的硬盘范围
	Volume int64 `name:"Volume"`
	// 切换访问新实例的方式，默认为0，升级主实例时，可指定该参数，升级只读实例或者灾备实例时指定该参数无意义，支持值包括：0-立刻切换，1-时间窗切换；当该值为1时，升级中过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口切换访问新实例触发该流程
	WaitSwitch *int64 `name:"WaitSwitch,omitempty"`
}

func (req *UpgradeDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeDbInstanceResponse, error) {
	resp := &UpgradeDbInstanceResponse{}
	err := client.Request("cdb", "UpgradeDBInstance", "2017-03-20").Do(req, resp)
	return resp, err
}

type UpgradeDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果
	AsyncRequestId string `json:"AsyncRequestId"`
	// 短订单ID，用于调用云API相关接口，如获取订单信息
	DealIds []*string `json:"DealIds"`
	// 长订单ID，用于反馈订单问题给腾讯云官方客服
	DealNames []*string `json:"DealNames"`
}

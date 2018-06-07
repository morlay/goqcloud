package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 升级实例版本
// https://cloud.tencent.com/document/api/236/15870

type UpgradeDbInstanceEngineVersionRequest struct {
	// 主实例数据库引擎版本，支持值包括：5.6和5.7
	EngineVersion string `name:"EngineVersion"`
	// 实例ID，格式如：cdb-c1nl9rpv或者cdbro-c1nl9rpv。与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 切换访问新实例的方式，默认为0，升级主实例时，可指定该参数，升级只读实例或者灾备实例时指定该参数无意义，支持值包括：0-立刻切换，1-时间窗切换；当该值为1时，升级中过程中，切换访问新实例的流程将会在时间窗内进行，或者用户主动调用接口切换访问新实例触发该流程
	WaitSwitch *int64 `name:"WaitSwitch,omitempty"`
}

func (req *UpgradeDbInstanceEngineVersionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpgradeDbInstanceEngineVersionResponse, error) {
	resp := &UpgradeDbInstanceEngineVersionResponse{}
	err := client.Request("cdb", "UpgradeDBInstanceEngineVersion", "2017-03-20").Do(req, resp)
	return resp, err
}

type UpgradeDbInstanceEngineVersionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务ID，可使用查询任务列表获取其执行情况
	AsyncRequestId string `json:"AsyncRequestId"`
}

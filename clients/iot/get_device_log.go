package iot

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取设备日志
// https://cloud.tencent.com/document/api/568/16523
type GetDeviceLogRequest struct {
	// 设备名称列表，最大支持100台
	DeviceNames []*string `name:"DeviceNames"`
	// 查询结束时间
	EndTime time.Time `name:"EndTime"`
	// 时间排序（desc/asc）
	Order *string `name:"Order,omitempty"`
	// 产品Id
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
	// 查询游标
	ScrollId *string `name:"ScrollId,omitempty"`
	// 查询数据量
	Size *int64 `name:"Size,omitempty"`
	// 查询开始时间
	StartTime time.Time `name:"StartTime"`
	// 日志类型（comm/status）
	Type *string `name:"Type,omitempty"`
}

func (req *GetDeviceLogRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetDeviceLogResponse, error) {
	resp := &GetDeviceLogResponse{}
	err := client.Request("iot", "GetDeviceLog", "2018-01-23").Do(req, resp)
	return resp, err
}

type GetDeviceLogResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备日志
	DeviceLog []*Object `json:"DeviceLog"`
	// 查询游标
	ScrollId []*string `json:"ScrollId"`
}

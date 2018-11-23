package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取机器详情
// https://cloud.tencent.com/document/api/296/19851

type DescribeMachineInfoRequest struct {
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeMachineInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeMachineInfoResponse, error) {
	resp := &DescribeMachineInfoResponse{}
	err := client.Request("yunjing", "DescribeMachineInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeMachineInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// CVM或BM主机唯一标识。
	InstanceId string `json:"InstanceId"`
	// 是否开通专业版。true：是false：否
	IsProVersion bool `json:"IsProVersion"`
	// 机器ip。
	MachineIp string `json:"MachineIp"`
	// 主机名称。
	MachineName string `json:"MachineName"`
	// 操作系统。
	MachineOs string `json:"MachineOs"`
	// 机器所属地域。如：ap-guangzhou，ap-shanghai
	MachineRegion string `json:"MachineRegion"`
	// 在线状态。ONLINE： 在线OFFLINE：离线
	MachineStatus string `json:"MachineStatus"`
	// 云主机类型。CVM: 虚拟主机BM: 黑石物理机
	MachineType string `json:"MachineType"`
	// 主机外网IP。
	MachineWanIp string `json:"MachineWanIp"`
	// 主机状态。POSTPAY: 表示后付费，即按量计费  PREPAY: 表示预付费，即包年包月
	PayMode string `json:"PayMode"`
	// 专业版开通时间。
	ProVersionOpenDate string `json:"ProVersionOpenDate"`
	// 受云镜保护天数。
	ProtectDays int64 `json:"ProtectDays"`
	// CVM或BM主机唯一Uuid。
	Quuid string `json:"Quuid"`
	// 云镜客户端唯一Uuid。
	Uuid string `json:"Uuid"`
}

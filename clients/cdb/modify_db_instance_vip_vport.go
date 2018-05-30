package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云数据库实例的IP和端口号
// https://cloud.tencent.com/document/api/236/15867
type ModifyDbInstanceVipVportRequest struct {
	// 目标IP。
	DstIp *string `name:"DstIp,omitempty"`
	// 目标端口，支持范围为：[1024-65535]。
	DstPort *int64 `name:"DstPort,omitempty"`
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 子网统一ID。
	UniqSubnetId *string `name:"UniqSubnetId,omitempty"`
	// 私有网络统一ID。
	UniqVpcId *string `name:"UniqVpcId,omitempty"`
}

func (req *ModifyDbInstanceVipVportRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDbInstanceVipVportResponse, error) {
	resp := &ModifyDbInstanceVipVportResponse{}
	err := client.Request("cdb", "ModifyDBInstanceVipVport", "2017-03-20").Do(req, resp)
	return resp, err
}

type ModifyDbInstanceVipVportResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务ID，可使用查询任务列表获取其执行情况。
	AsyncRequestId string `json:"AsyncRequestId"`
}

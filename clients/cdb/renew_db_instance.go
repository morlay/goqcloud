package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费云数据库实例
// https://cloud.tencent.com/document/api/236/30160

type RenewDbInstanceRequest struct {
	// 待续费的实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
	// 续费时长，单位：月，可选值包括[1,2,3,4,5,6,7,8,9,10,11,12,24,36]
	TimeSpan int64 `name:"TimeSpan"`
}

func (req *RenewDbInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewDbInstanceResponse, error) {
	resp := &RenewDbInstanceResponse{}
	err := client.Request("cdb", "RenewDBInstance", "2017-03-20").Do(req, resp)
	return resp, err
}

type RenewDbInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 订单ID
	DealId string `json:"DealId"`
}

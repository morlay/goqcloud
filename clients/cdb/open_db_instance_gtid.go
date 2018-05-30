package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 开启实例的GTID
// https://cloud.tencent.com/document/api/236/17489
type OpenDbInstanceGtidRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *OpenDbInstanceGtidRequest) Invoke(client github_com_morlay_goqcloud.Client) (*OpenDbInstanceGtidResponse, error) {
	resp := &OpenDbInstanceGtidResponse{}
	err := client.Request("cdb", "OpenDBInstanceGTID", "2017-03-20").Do(req, resp)
	return resp, err
}

type OpenDbInstanceGtidResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}

package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 开通实例外网访问
// https://cloud.tencent.com/document/api/236/15875

type OpenWanServiceRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *OpenWanServiceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*OpenWanServiceResponse, error) {
	resp := &OpenWanServiceResponse{}
	err := client.Request("cdb", "OpenWanService", "2017-03-20").Do(req, resp)
	return resp, err
}

type OpenWanServiceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果
	AsyncRequestId string `json:"AsyncRequestId"`
}

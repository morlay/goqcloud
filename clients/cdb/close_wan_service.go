package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 关闭实例外网访问
// https://cloud.tencent.com/document/api/236/15863
type CloseWanServiceRequest struct {
	// 实例ID，格式如：cdb-c1nl9rpv，与云数据库控制台页面中显示的实例ID相同，可使用查询实例列表 接口获取，其值为输出参数中字段 InstanceId 的值。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CloseWanServiceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CloseWanServiceResponse, error) {
	resp := &CloseWanServiceResponse{}
	err := client.Request("cdb", "CloseWanService", "2017-03-20").Do(req, resp)
	return resp, err
}

type CloseWanServiceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务的请求ID，可使用此ID查询异步任务的执行结果。
	AsyncRequestId string `json:"AsyncRequestId"`
}

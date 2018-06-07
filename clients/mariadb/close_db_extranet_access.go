package mariadb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 关闭外网访问
// https://cloud.tencent.com/document/api/237/16179

type CloseDbExtranetAccessRequest struct {
	// 待关闭外网访问的实例ID。形如：tdsql-ow728lmc，可以通过 DescribeDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CloseDbExtranetAccessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CloseDbExtranetAccessResponse, error) {
	resp := &CloseDbExtranetAccessResponse{}
	err := client.Request("mariadb", "CloseDBExtranetAccess", "2017-03-12").Do(req, resp)
	return resp, err
}

type CloseDbExtranetAccessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowId int64 `json:"FlowId"`
}

package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 关闭外网访问
// https://cloud.tencent.com/document/api/557/19998

type CloseDbExtranetAccessRequest struct {
	// 待关闭外网访问的实例ID。形如：dcdbt-ow728lmc，可以通过 DescribeDCDBInstances 查询实例详情获得。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CloseDbExtranetAccessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CloseDbExtranetAccessResponse, error) {
	resp := &CloseDbExtranetAccessResponse{}
	err := client.Request("dcdb", "CloseDBExtranetAccess", "2018-04-11").Do(req, resp)
	return resp, err
}

type CloseDbExtranetAccessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowId int64 `json:"FlowId"`
}

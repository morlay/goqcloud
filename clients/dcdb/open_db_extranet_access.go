package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 开通外网访问
// https://cloud.tencent.com/document/api/557/19980

type OpenDbExtranetAccessRequest struct {
	// 待开放外网访问的实例ID。形如：dcdbt-ow728lmc。
	InstanceId string `name:"InstanceId"`
	// 区域
	Region string `name:"Region"`
}

func (req *OpenDbExtranetAccessRequest) Invoke(client github_com_morlay_goqcloud.Client) (*OpenDbExtranetAccessResponse, error) {
	resp := &OpenDbExtranetAccessResponse{}
	err := client.Request("dcdb", "OpenDBExtranetAccess", "2018-04-11").Do(req, resp)
	return resp, err
}

type OpenDbExtranetAccessResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 异步任务Id，可通过 DescribeFlow 查询任务状态。
	FlowId int64 `json:"FlowId"`
}

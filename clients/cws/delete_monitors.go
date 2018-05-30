package cws

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除监控任务
// https://cloud.tencent.com/document/api/692/16744
type DeleteMonitorsRequest struct {
	// 监控任务ID列表
	MonitorIds []*int64 `name:"MonitorIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteMonitorsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteMonitorsResponse, error) {
	resp := &DeleteMonitorsResponse{}
	err := client.Request("cws", "DeleteMonitors", "2018-03-12").Do(req, resp)
	return resp, err
}

type DeleteMonitorsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

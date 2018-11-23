package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询日志
// https://cloud.tencent.com/document/api/851/18310

type QueryLogsRequest struct {
	// 集群名称
	Cluster string `name:"Cluster"`
	// 加载更多使用，透传上次返回的context值，获取后续的日志内容，使用context翻页最多能获取10000条日志
	Context *string `name:"Context,omitempty"`
	// 查询日志的结束时间
	EndTime string `name:"EndTime"`
	// 任务名称
	JobName string `name:"JobName"`
	// 单次要返回的日志条数
	Limit int64 `name:"Limit"`
	// 区域
	Region string `name:"Region"`
	// 查询日志的开始时间
	StartTime string `name:"StartTime"`
}

func (req *QueryLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*QueryLogsResponse, error) {
	resp := &QueryLogsResponse{}
	err := client.Request("tia", "QueryLogs", "2018-02-26").Do(req, resp)
	return resp, err
}

type QueryLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 日志查询上下文，用于加载更多日志
	Context string `json:"Context"`
	// 是否已经返回所有符合条件的日志
	Listover bool `json:"Listover"`
	// 日志内容列表
	Logs []*Log `json:"Logs"`
}

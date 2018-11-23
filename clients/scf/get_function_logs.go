package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取函数运行日志
// https://cloud.tencent.com/document/api/583/18583

type GetFunctionLogsRequest struct {
	// 查询的具体日期，例如：2017-05-16 20:59:59，只能与startTime相差一天之内
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 日志过滤条件。可用来区分正确和错误日志，filter.retCode=not0 表示只返回错误日志，filter.retCode=is0 表示只返回正确日志，不传，则返回所有日志
	Filter *Filter `name:"Filter,omitempty"`
	// 函数的名称
	FunctionName *string `name:"FunctionName,omitempty"`
	// 执行该函数对应的requestId
	FunctionRequestId *string `name:"FunctionRequestId,omitempty"`
	// 返回数据的长度，Offset+Limit不能大于10000
	Limit *int64 `name:"Limit,omitempty"`
	// 数据的偏移量，Offset+Limit不能大于10000
	Offset *int64 `name:"Offset,omitempty"`
	// 以升序还是降序的方式对日志进行排序，可选值 desc和 acs
	Order *string `name:"Order,omitempty"`
	// 根据某个字段排序日志,支持以下字段：startTime、functionName、requestId、duration和 memUsage
	OrderBy *string `name:"OrderBy,omitempty"`
	// 函数的版本
	Qualifier *string `name:"Qualifier,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 查询的具体日期，例如：2017-05-16 20:00:00，只能与endtime相差一天之内
	StartTime *time.Time `name:"StartTime,omitempty"`
}

func (req *GetFunctionLogsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetFunctionLogsResponse, error) {
	resp := &GetFunctionLogsResponse{}
	err := client.Request("scf", "GetFunctionLogs", "2018-04-16").Do(req, resp)
	return resp, err
}

type GetFunctionLogsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 函数日志信息
	Data []*FunctionLog `json:"Data"`
	// 函数日志的总数
	TotalCount int64 `json:"TotalCount"`
}

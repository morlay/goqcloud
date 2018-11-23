package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取进程列表
// https://cloud.tencent.com/document/api/296/30332

type DescribeProcessesRequest struct {
	// 过滤条件。ProcessName - String - 是否必填：否 - 进程名MachineIp - String - 是否必填：否 - 主机内网IP
	Filters []*Filter `name:"Filters,omitempty"`
	// 返回数量，默认为10，最大值为100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 进程名。Uuid和ProcessName必填其一，使用ProcessName表示，查询该进程列表信息。
	ProcessName *string `name:"ProcessName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 云镜客户端唯一Uuid。Uuid和ProcessName必填其一，使用Uuid表示，查询该主机列表信息。
	Uuid *string `name:"Uuid,omitempty"`
}

func (req *DescribeProcessesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeProcessesResponse, error) {
	resp := &DescribeProcessesResponse{}
	err := client.Request("yunjing", "DescribeProcesses", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeProcessesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 进程列表数据数组。
	Processes []*Process `json:"Processes"`
	// 进程列表记录总数。
	TotalCount int64 `json:"TotalCount"`
}

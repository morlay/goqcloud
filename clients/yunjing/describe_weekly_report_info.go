package yunjing

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取专业周报详情
// https://cloud.tencent.com/document/api/296/30325

type DescribeWeeklyReportInfoRequest struct {
	// 专业周报开始时间。
	BeginDate time.Time `name:"BeginDate"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeWeeklyReportInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeWeeklyReportInfoResponse, error) {
	resp := &DescribeWeeklyReportInfoResponse{}
	err := client.Request("yunjing", "DescribeWeeklyReportInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeWeeklyReportInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 周报开始时间。
	BeginDate time.Time `json:"BeginDate"`
	// 密码破解成功数。
	BruteAttackSuccessNum int64 `json:"BruteAttackSuccessNum"`
	// 账号所属公司或个人名称。
	CompanyName string `json:"CompanyName"`
	// 周报结束时间。
	EndDate time.Time `json:"EndDate"`
	// 安全等级。HIGH：高MIDDLE：中LOW：低
	Level string `json:"Level"`
	// 机器总数。
	MachineNum int64 `json:"MachineNum"`
	// 木马记录数。
	MalwareNum int64 `json:"MalwareNum"`
	// 异地登录数。
	NonlocalLoginNum int64 `json:"NonlocalLoginNum"`
	// 云镜客户端离线数。
	OfflineMachineNum int64 `json:"OfflineMachineNum"`
	// 云镜客户端在线数。
	OnlineMachineNum int64 `json:"OnlineMachineNum"`
	// 开通云镜专业版数量。
	ProVersionMachineNum int64 `json:"ProVersionMachineNum"`
	// 漏洞数。
	VulNum int64 `json:"VulNum"`
}

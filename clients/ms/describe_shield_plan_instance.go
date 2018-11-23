package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询加固策略
// https://cloud.tencent.com/document/api/283/18575

type DescribeShieldPlanInstanceRequest struct {
	// 服务类别id
	Pid int64 `name:"Pid"`
	// 区域
	Region string `name:"Region"`
	// 资源id
	ResourceId string `name:"ResourceId"`
}

func (req *DescribeShieldPlanInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeShieldPlanInstanceResponse, error) {
	resp := &DescribeShieldPlanInstanceResponse{}
	err := client.Request("ms", "DescribeShieldPlanInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type DescribeShieldPlanInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 绑定资源信息
	BindInfo BindInfo `json:"BindInfo"`
	// 加固资源信息
	ResourceServiceInfo ResourceServiceInfo `json:"ResourceServiceInfo"`
	// 加固策略信息
	ShieldPlanInfo ShieldPlanInfo `json:"ShieldPlanInfo"`
}

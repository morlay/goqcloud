package ms

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 新增加固策略
// https://cloud.tencent.com/document/api/283/18576

type CreateShieldPlanInstanceRequest struct {
	// 策略具体信息
	PlanInfo PlanInfo `name:"PlanInfo"`
	// 策略名称
	PlanName string `name:"PlanName"`
	// 区域
	Region string `name:"Region"`
	// 资源id
	ResourceId string `name:"ResourceId"`
}

func (req *CreateShieldPlanInstanceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateShieldPlanInstanceResponse, error) {
	resp := &CreateShieldPlanInstanceResponse{}
	err := client.Request("ms", "CreateShieldPlanInstance", "2018-04-08").Do(req, resp)
	return resp, err
}

type CreateShieldPlanInstanceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 策略id
	PlanId int64 `json:"PlanId"`
	// 任务状态: 1-已完成,2-处理中,3-处理出错,4-处理超时
	Progress int64 `json:"Progress"`
}

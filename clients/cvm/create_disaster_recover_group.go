package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 创建分散置放群组
// https://cloud.tencent.com/document/api/213/17813

type CreateDisasterRecoverGroupRequest struct {
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `name:"ClientToken,omitempty"`
	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
	// 分散置放群组类型，取值范围：HOST：物理机SW：交换机RACK：机架
	Type string `name:"Type"`
}

func (req *CreateDisasterRecoverGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDisasterRecoverGroupResponse, error) {
	resp := &CreateDisasterRecoverGroupResponse{}
	err := client.Request("cvm", "CreateDisasterRecoverGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDisasterRecoverGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 置放群组创建时间。
	CreateTime time.Time `json:"CreateTime"`
	// 置放群组内已有的云主机数量。
	CurrentNum int64 `json:"CurrentNum"`
	// 置放群组内可容纳的云主机数量。
	CvmQuotaTotal int64 `json:"CvmQuotaTotal"`
	// 分散置放群组ID列表。
	DisasterRecoverGroupId string `json:"DisasterRecoverGroupId"`
	// 分散置放群组名称，长度1-60个字符，支持中、英文。
	Name string `json:"Name"`
	// 分散置放群组类型，取值范围：HOST：物理机SW：交换机RACK：机架
	Type string `json:"Type"`
}

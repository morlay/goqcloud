package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改实例计费模式
// https://cloud.tencent.com/document/api/213/17964

type ModifyInstancesChargeTypeRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `name:"InstanceChargePrepaid,omitempty"`
	// 实例计费类型。PREPAID：预付费，即包年包月。
	InstanceChargeType string `name:"InstanceChargeType"`
	// 一个或多个待操作的实例ID。可通过DescribeInstances接口返回值中的InstanceId获取。每次请求批量实例的上限为100。
	InstanceIds []*string `name:"InstanceIds"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyInstancesChargeTypeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyInstancesChargeTypeResponse, error) {
	resp := &ModifyInstancesChargeTypeResponse{}
	err := client.Request("cvm", "ModifyInstancesChargeType", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyInstancesChargeTypeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

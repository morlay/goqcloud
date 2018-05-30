package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 安全组添加规则
// https://cloud.tencent.com/document/api/215/15807
type CreateSecurityGroupPoliciesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
	// 安全组规则集合。
	SecurityGroupPolicySet SecurityGroupPolicySet `name:"SecurityGroupPolicySet"`
}

func (req *CreateSecurityGroupPoliciesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSecurityGroupPoliciesResponse, error) {
	resp := &CreateSecurityGroupPoliciesResponse{}
	err := client.Request("vpc", "CreateSecurityGroupPolicies", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateSecurityGroupPoliciesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

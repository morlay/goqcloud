package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询安全组规则
// https://cloud.tencent.com/document/api/215/15804

type DescribeSecurityGroupPoliciesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如：sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
}

func (req *DescribeSecurityGroupPoliciesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSecurityGroupPoliciesResponse, error) {
	resp := &DescribeSecurityGroupPoliciesResponse{}
	err := client.Request("vpc", "DescribeSecurityGroupPolicies", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupPoliciesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全组规则集合。
	SecurityGroupPolicySet SecurityGroupPolicySet `json:"SecurityGroupPolicySet"`
}

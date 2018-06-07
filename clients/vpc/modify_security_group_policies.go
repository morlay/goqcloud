package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改安全组出站和入站规则
// https://cloud.tencent.com/document/api/215/15810

type ModifySecurityGroupPoliciesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
	// 安全组规则集合。 SecurityGroupPolicySet对象必须同时指定新的出（Egress）入（Ingress）站规则。 SecurityGroupPolicy对象不支持自定义索引（PolicyIndex）。
	SecurityGroupPolicySet SecurityGroupPolicySet `name:"SecurityGroupPolicySet"`
}

func (req *ModifySecurityGroupPoliciesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifySecurityGroupPoliciesResponse, error) {
	resp := &ModifySecurityGroupPoliciesResponse{}
	err := client.Request("vpc", "ModifySecurityGroupPolicies", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifySecurityGroupPoliciesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

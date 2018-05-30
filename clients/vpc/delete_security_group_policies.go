package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除安全组规则
// https://cloud.tencent.com/document/api/215/15809
type DeleteSecurityGroupPoliciesRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
	// 安全组规则集合。一个请求中只能删除单个方向的一条或多条规则。支持指定索引（PolicyIndex） 匹配删除和安全组规则匹配删除两种方式，一个请求中只能使用一种匹配方式。
	SecurityGroupPolicySet SecurityGroupPolicySet `name:"SecurityGroupPolicySet"`
}

func (req *DeleteSecurityGroupPoliciesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSecurityGroupPoliciesResponse, error) {
	resp := &DeleteSecurityGroupPoliciesResponse{}
	err := client.Request("vpc", "DeleteSecurityGroupPolicies", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteSecurityGroupPoliciesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

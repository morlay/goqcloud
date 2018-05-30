package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 替换单条安全组路由规则
// https://cloud.tencent.com/document/api/215/15811
type ReplaceSecurityGroupPolicyRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
	// 安全组规则集合对象。
	SecurityGroupPolicySet SecurityGroupPolicySet `name:"SecurityGroupPolicySet"`
}

func (req *ReplaceSecurityGroupPolicyRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ReplaceSecurityGroupPolicyResponse, error) {
	resp := &ReplaceSecurityGroupPolicyResponse{}
	err := client.Request("vpc", "ReplaceSecurityGroupPolicy", "2017-03-12").Do(req, resp)
	return resp, err
}

type ReplaceSecurityGroupPolicyResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

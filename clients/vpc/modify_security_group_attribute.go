package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改安全组属性
// https://cloud.tencent.com/document/api/215/15805

type ModifySecurityGroupAttributeRequest struct {
	// 安全组备注，最多100个字符。
	GroupDescription *string `name:"GroupDescription,omitempty"`
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName *string `name:"GroupName,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
}

func (req *ModifySecurityGroupAttributeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifySecurityGroupAttributeResponse, error) {
	resp := &ModifySecurityGroupAttributeResponse{}
	err := client.Request("vpc", "ModifySecurityGroupAttribute", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifySecurityGroupAttributeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

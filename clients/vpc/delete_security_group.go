package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除安全组
// https://cloud.tencent.com/document/api/215/15803

type DeleteSecurityGroupRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全组实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupId string `name:"SecurityGroupId"`
}

func (req *DeleteSecurityGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteSecurityGroupResponse, error) {
	resp := &DeleteSecurityGroupResponse{}
	err := client.Request("vpc", "DeleteSecurityGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type DeleteSecurityGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

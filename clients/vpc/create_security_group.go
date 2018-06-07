package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建安全组
// https://cloud.tencent.com/document/api/215/15806

type CreateSecurityGroupRequest struct {
	// 安全组备注，最多100个字符。
	GroupDescription string `name:"GroupDescription"`
	// 安全组名称，可任意命名，但不得超过60个字符。
	GroupName string `name:"GroupName"`
	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateSecurityGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSecurityGroupResponse, error) {
	resp := &CreateSecurityGroupResponse{}
	err := client.Request("vpc", "CreateSecurityGroup", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateSecurityGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全组对象。
	SecurityGroup SecurityGroup `json:"SecurityGroup"`
}

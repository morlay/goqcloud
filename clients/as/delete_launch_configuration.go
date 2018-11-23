package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除启动配置
// https://cloud.tencent.com/document/api/377/20446

type DeleteLaunchConfigurationRequest struct {
	// 需要删除的启动配置ID。
	LaunchConfigurationId string `name:"LaunchConfigurationId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteLaunchConfigurationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteLaunchConfigurationResponse, error) {
	resp := &DeleteLaunchConfigurationResponse{}
	err := client.Request("as", "DeleteLaunchConfiguration", "2018-04-19").Do(req, resp)
	return resp, err
}

type DeleteLaunchConfigurationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

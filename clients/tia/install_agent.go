package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 安装agent
// https://cloud.tencent.com/document/api/851/18303

type InstallAgentRequest struct {
	// 集群名称
	Cluster string `name:"Cluster"`
	// 区域
	Region string `name:"Region"`
	// Agent版本, 用于私有集群的agent安装，默认为“private-training”
	TiaVersion *string `name:"TiaVersion,omitempty"`
	// 是否允许更新Agent
	Update *bool `name:"Update,omitempty"`
}

func (req *InstallAgentRequest) Invoke(client github_com_morlay_goqcloud.Client) (*InstallAgentResponse, error) {
	resp := &InstallAgentResponse{}
	err := client.Request("tia", "InstallAgent", "2018-02-26").Do(req, resp)
	return resp, err
}

type InstallAgentResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// Agent版本, 用于私有集群的agent安装
	TiaVersion string `json:"TiaVersion"`
}

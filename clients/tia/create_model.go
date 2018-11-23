package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建Model
// https://cloud.tencent.com/document/api/851/18315

type CreateModelRequest struct {
	// 指定集群的名称（集群模式下必填）
	Cluster *string `name:"Cluster,omitempty"`
	// 关于模型的描述
	Description *string `name:"Description,omitempty"`
	// 暴露外网或内网，默认暴露外网（集群模式下选填）
	Expose *string `name:"Expose,omitempty"`
	// 要部署模型的路径名
	Model string `name:"Model"`
	// 模型名称
	Name string `name:"Name"`
	// 区域
	Region string `name:"Region"`
	// 要部署的模型副本数目（集群模式下选填）
	Replicas *int64 `name:"Replicas,omitempty"`
	// 部署模型的其他配置信息
	RuntimeConf []*string `name:"RuntimeConf,omitempty"`
	// 运行环境镜像的标签
	RuntimeVersion *string `name:"RuntimeVersion,omitempty"`
	// 部署模式（无服务器函数模式/集群模式）
	ServType *string `name:"ServType,omitempty"`
}

func (req *CreateModelRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateModelResponse, error) {
	resp := &CreateModelResponse{}
	err := client.Request("tia", "CreateModel", "2018-02-26").Do(req, resp)
	return resp, err
}

type CreateModelResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 模型的详细信息
	Model Model `json:"Model"`
}

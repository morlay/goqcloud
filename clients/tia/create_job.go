package tia

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建Job
// https://cloud.tencent.com/document/api/851/18308

type CreateJobRequest struct {
	// 任务启动参数
	Args []*string `name:"Args,omitempty"`
	// 运行任务的集群
	Cluster string `name:"Cluster"`
	// 任务启动命令
	Command []*string `name:"Command,omitempty"`
	// 启动debug mode，默认为false
	Debug *bool `name:"Debug,omitempty"`
	// （ScaleTier为Custom时）master机器类型
	MasterType *string `name:"MasterType,omitempty"`
	// 任务名称
	Name string `name:"Name"`
	// 挂载的路径，支持nfs,cos(cos只在tia运行环境中支持)
	PackageDir []*string `name:"PackageDir,omitempty"`
	// （ScaleTier为Custom时）parameter server机器数量
	ParameterServerCount *int64 `name:"ParameterServerCount,omitempty"`
	// （ScaleTier为Custom时）parameter server机器类型
	ParameterServerType *string `name:"ParameterServerType,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 运行任务的其他配置信息
	RuntimeConf []*string `name:"RuntimeConf,omitempty"`
	// 运行任务的环境
	RuntimeVersion string `name:"RuntimeVersion"`
	// 运行任务的配置信息
	ScaleTier *string `name:"ScaleTier,omitempty"`
	// （ScaleTier为Custom时）worker机器数量
	WorkerCount *int64 `name:"WorkerCount,omitempty"`
	// （ScaleTier为Custom时）worker机器类型
	WorkerType *string `name:"WorkerType,omitempty"`
}

func (req *CreateJobRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateJobResponse, error) {
	resp := &CreateJobResponse{}
	err := client.Request("tia", "CreateJob", "2018-02-26").Do(req, resp)
	return resp, err
}

type CreateJobResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 训练任务信息
	Job Job `json:"Job"`
}

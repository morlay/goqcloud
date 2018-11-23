package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新函数配置
// https://cloud.tencent.com/document/api/583/18580

type UpdateFunctionConfigurationRequest struct {
	// 函数描述。最大支持 1000 个英文字母、数字、空格、逗号和英文句号，支持中文
	Description *string `name:"Description,omitempty"`
	// 函数的环境变量
	Environment *Environment `name:"Environment,omitempty"`
	// 要修改的函数名称
	FunctionName string `name:"FunctionName"`
	// 函数运行时内存大小，默认为 128 M，可选范 128 M-1536 M
	MemorySize *int64 `name:"MemorySize,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 函数运行环境，目前仅支持 Python2.7，Python3.6，Nodejs6.10，PHP5， PHP7，Golang1 和 Java8
	Runtime *string `name:"Runtime,omitempty"`
	// 函数最长执行时间，单位为秒，可选值范 1-300 秒，默认为 3 秒
	Timeout *int64 `name:"Timeout,omitempty"`
	// 函数的私有网络配置
	VpcConfig *VpcConfig `name:"VpcConfig,omitempty"`
}

func (req *UpdateFunctionConfigurationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateFunctionConfigurationResponse, error) {
	resp := &UpdateFunctionConfigurationResponse{}
	err := client.Request("scf", "UpdateFunctionConfiguration", "2018-04-16").Do(req, resp)
	return resp, err
}

type UpdateFunctionConfigurationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

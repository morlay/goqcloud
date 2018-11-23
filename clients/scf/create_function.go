package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建函数
// https://cloud.tencent.com/document/api/583/18586

type CreateFunctionRequest struct {
	// 函数的代码. 注意：不能同时指定Cos与ZipFile
	Code Code `name:"Code"`
	// 函数描述,最大支持 1000 个英文字母、数字、空格、逗号、换行符和英文句号，支持中文
	Description *string `name:"Description,omitempty"`
	// 函数的环境变量
	Environment *Environment `name:"Environment,omitempty"`
	// 创建的函数名称，函数名称支持26个英文字母大小写、数字、连接符和下划线，第一个字符只能以字母开头，最后一个字符不能为连接符或者下划线，名称长度2-60
	FunctionName string `name:"FunctionName"`
	// 函数处理方法名称，名称格式支持 "文件名称.方法名称" 形式，文件名称和函数名称之间以"."隔开，文件名称和函数名称要求以字母开始和结尾，中间允许插入字母、数字、下划线和连接符，文件名称和函数名字的长度要求是 2-60 个字符
	Handler *string `name:"Handler,omitempty"`
	// 函数运行时内存大小，默认为 128M，可选范围 128MB-1536MB，并且以 128MB 为阶梯
	MemorySize *int64 `name:"MemorySize,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 函数运行环境，目前仅支持 Python2.7，Python3.6，Nodejs6.10， PHP5， PHP7，Golang1 和 Java8，默认Python2.7
	Runtime *string `name:"Runtime,omitempty"`
	// 函数最长执行时间，单位为秒，可选值范围 1-300 秒，默认为 3 秒
	Timeout *int64 `name:"Timeout,omitempty"`
	// 函数的私有网络配置
	VpcConfig *VpcConfig `name:"VpcConfig,omitempty"`
}

func (req *CreateFunctionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateFunctionResponse, error) {
	resp := &CreateFunctionResponse{}
	err := client.Request("scf", "CreateFunction", "2018-04-16").Do(req, resp)
	return resp, err
}

type CreateFunctionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

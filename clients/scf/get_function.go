package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 获取函数详细信息
// https://cloud.tencent.com/document/api/583/18584

type GetFunctionRequest struct {
	// 需要获取详情的函数名称
	FunctionName string `name:"FunctionName"`
	// 函数的版本号
	Qualifier *string `name:"Qualifier,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 是否显示代码, TRUE表示显示代码，FALSE表示不显示代码,大于1M的入口文件不会显示
	ShowCode *string `name:"ShowCode,omitempty"`
}

func (req *GetFunctionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*GetFunctionResponse, error) {
	resp := &GetFunctionResponse{}
	err := client.Request("scf", "GetFunction", "2018-04-16").Do(req, resp)
	return resp, err
}

type GetFunctionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 代码错误信息
	CodeError string `json:"CodeError"`
	// 函数的代码
	CodeInfo string `json:"CodeInfo"`
	// 代码是否正确
	CodeResult string `json:"CodeResult"`
	// 函数代码大小
	CodeSize int64 `json:"CodeSize"`
	// 函数的描述信息
	Description string `json:"Description"`
	// 函数的环境变量
	Environment Environment `json:"Environment"`
	// 代码错误码
	ErrNo int64 `json:"ErrNo"`
	// 函数的名称
	FunctionName string `json:"FunctionName"`
	// 函数的版本
	FunctionVersion string `json:"FunctionVersion"`
	// 函数的入口
	Handler string `json:"Handler"`
	// 函数的最大可用内存
	MemorySize int64 `json:"MemorySize"`
	// 函数的最后修改时间
	ModTime time.Time `json:"ModTime"`
	// 函数的命名空间
	Namespace string `json:"Namespace"`
	// 函数绑定的角色
	Role string `json:"Role"`
	// 函数的运行环境
	Runtime string `json:"Runtime"`
	// 函数的超时时间
	Timeout int64 `json:"Timeout"`
	// 函数的触发器列表
	Triggers []*Trigger `json:"Triggers"`
	// 是否使用GPU
	UseGpu string `json:"UseGpu"`
	// 函数的私有网络
	VpcConfig VpcConfig `json:"VpcConfig"`
}

package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新函数代码
// https://cloud.tencent.com/document/api/583/18581

type UpdateFunctionCodeRequest struct {
	// 对象存储桶名称
	CosBucketName *string `name:"CosBucketName,omitempty"`
	// 对象存储的地域，地域为北京时需要传入ap-beijing,北京一区时需要传递ap-beijing-1，其他的地域不需要传递。
	CosBucketRegion *string `name:"CosBucketRegion,omitempty"`
	// 对象存储对象路径
	CosObjectName *string `name:"CosObjectName,omitempty"`
	// 要修改的函数名称
	FunctionName string `name:"FunctionName"`
	// 函数处理方法名称。名称格式支持“文件名称.函数名称”形式，文件名称和函数名称之间以"."隔开，文件名称和函数名称要求以字母开始和结尾，中间允许插入字母、数字、下划线和连接符，文件名称和函数名字的长度要求 2-60 个字符
	Handler string `name:"Handler"`
	// 区域
	Region string `name:"Region"`
	// 包含函数代码文件及其依赖项的 zip 格式文件，使用该接口时要求将 zip 文件的内容转成 base64 编码，最大支持20M
	ZipFile *string `name:"ZipFile,omitempty"`
}

func (req *UpdateFunctionCodeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateFunctionCodeResponse, error) {
	resp := &UpdateFunctionCodeResponse{}
	err := client.Request("scf", "UpdateFunctionCode", "2018-04-16").Do(req, resp)
	return resp, err
}

type UpdateFunctionCodeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

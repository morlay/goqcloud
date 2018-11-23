package scf

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除函数
// https://cloud.tencent.com/document/api/583/18585

type DeleteFunctionRequest struct {
	// 要删除的函数名称
	FunctionName string `name:"FunctionName"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeleteFunctionRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteFunctionResponse, error) {
	resp := &DeleteFunctionResponse{}
	err := client.Request("scf", "DeleteFunction", "2018-04-16").Do(req, resp)
	return resp, err
}

type DeleteFunctionResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

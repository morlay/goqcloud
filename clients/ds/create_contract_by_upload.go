package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 通过上传创建合同
// https://cloud.tencent.com/document/api/869/17795

type CreateContractByUploadRequest struct {
	// 合同上传链接地址
	ContractFile string `name:"ContractFile"`
	// 合同名称
	ContractName string `name:"ContractName"`
	// 合同发起方帐号ID
	Initiator string `name:"Initiator"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
	// 备注
	Remarks *string `name:"Remarks,omitempty"`
	// 签署人信息
	SignInfos []*SignInfo `name:"SignInfos"`
}

func (req *CreateContractByUploadRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateContractByUploadResponse, error) {
	resp := &CreateContractByUploadResponse{}
	err := client.Request("ds", "CreateContractByUpload", "2018-05-23").Do(req, resp)
	return resp, err
}

type CreateContractByUploadResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	TaskId int64 `json:"TaskId"`
}

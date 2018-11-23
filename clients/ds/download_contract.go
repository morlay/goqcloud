package ds

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 下载合同
// https://cloud.tencent.com/document/api/869/17788

type DownloadContractRequest struct {
	// 合同ID
	ContractResId string `name:"ContractResId"`
	// 模块名
	Module string `name:"Module"`
	// 操作名
	Operation string `name:"Operation"`
	// 区域
	Region string `name:"Region"`
}

func (req *DownloadContractRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DownloadContractResponse, error) {
	resp := &DownloadContractResponse{}
	err := client.Request("ds", "DownloadContract", "2018-05-23").Do(req, resp)
	return resp, err
}

type DownloadContractResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	TaskId int64 `json:"TaskId"`
}

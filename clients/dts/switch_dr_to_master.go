package dts

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 灾备升级为主实例
// https://cloud.tencent.com/document/api/571/18567

type SwitchDrToMasterRequest struct {
	// 数据库的类型  （如 mysql）
	DatabaseType string `name:"DatabaseType"`
	// 灾备实例的信息
	DstInfo SyncInstanceInfo `name:"DstInfo"`
	// 区域
	Region string `name:"Region"`
}

func (req *SwitchDrToMasterRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SwitchDrToMasterResponse, error) {
	resp := &SwitchDrToMasterResponse{}
	err := client.Request("dts", "SwitchDrToMaster", "2018-03-30").Do(req, resp)
	return resp, err
}

type SwitchDrToMasterResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 后台异步任务请求id
	AsyncRequestId string `json:"AsyncRequestId"`
}

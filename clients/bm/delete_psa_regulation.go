package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除预授权规则
// https://cloud.tencent.com/document/api/386/30142

type DeletePsaRegulationRequest struct {
	// 预授权规则ID
	PsaId string `name:"PsaId"`
	// 区域
	Region string `name:"Region"`
}

func (req *DeletePsaRegulationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeletePsaRegulationResponse, error) {
	resp := &DeletePsaRegulationResponse{}
	err := client.Request("bm", "DeletePsaRegulation", "2018-04-23").Do(req, resp)
	return resp, err
}

type DeletePsaRegulationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

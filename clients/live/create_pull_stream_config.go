package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 添加拉流配置
// https://cloud.tencent.com/document/api/267/30159

type CreatePullStreamConfigRequest struct {
	// 区域id,1-深圳,2-上海，3-天津,4-香港。
	AreaId int64 `name:"AreaId"`
	// 结束时间，注意：1. 结束时间必须大于开始时间；2. 结束时间和开始时间必须大于当前时间；3. 结束时间 和 开始时间 间隔必须小于七天。
	EndTime time.Time `name:"EndTime"`
	// 源Url。
	FromUrl string `name:"FromUrl"`
	// 运营商id,1-电信,2-移动,3-联通,4-其他,AreaId为4的时候,IspId只能为其他。
	IspId int64 `name:"IspId"`
	// 区域
	Region string `name:"Region"`
	// 开始时间。
	StartTime time.Time `name:"StartTime"`
	// 目的Url，目前限制该目标地址为腾讯域名。
	ToUrl string `name:"ToUrl"`
}

func (req *CreatePullStreamConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreatePullStreamConfigResponse, error) {
	resp := &CreatePullStreamConfigResponse{}
	err := client.Request("live", "CreatePullStreamConfig", "2018-08-01").Do(req, resp)
	return resp, err
}

type CreatePullStreamConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 配置成功后的id。
	ConfigId string `json:"ConfigId"`
}

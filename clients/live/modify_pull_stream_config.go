package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
	time "time"
)

// 更新拉流配置
// https://cloud.tencent.com/document/api/267/30157

type ModifyPullStreamConfigRequest struct {
	// 区域id,1-深圳,2-上海，3-天津,4-香港。如有改动，需同时传入IspId。
	AreaId *int64 `name:"AreaId,omitempty"`
	// 配置id。
	ConfigId string `name:"ConfigId"`
	// 结束时间，注意：1. 结束时间必须大于开始时间；2. 结束时间和开始时间必须大于当前时间；3. 结束时间 和 开始时间 间隔必须小于七天。
	EndTime *time.Time `name:"EndTime,omitempty"`
	// 源Url。
	FromUrl *string `name:"FromUrl,omitempty"`
	// 运营商id,1-电信,2-移动,3-联通,4-其他,AreaId为4的时候,IspId只能为其他。如有改动，需同时传入AreaId。
	IspId *int64 `name:"IspId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 开始时间。
	StartTime *time.Time `name:"StartTime,omitempty"`
	// 目的Url。
	ToUrl *string `name:"ToUrl,omitempty"`
}

func (req *ModifyPullStreamConfigRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyPullStreamConfigResponse, error) {
	resp := &ModifyPullStreamConfigResponse{}
	err := client.Request("live", "ModifyPullStreamConfig", "2018-08-01").Do(req, resp)
	return resp, err
}

type ModifyPullStreamConfigResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

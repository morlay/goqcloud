package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 设置水印状态
// https://cloud.tencent.com/document/api/267/30151

type SetLiveWatermarkStatusRequest struct {
	// 区域
	Region string `name:"Region"`
	// 状态。0：停用，1:启用
	Status int64 `name:"Status"`
	// 水印ID。
	WatermarkId int64 `name:"WatermarkId"`
}

func (req *SetLiveWatermarkStatusRequest) Invoke(client github_com_morlay_goqcloud.Client) (*SetLiveWatermarkStatusResponse, error) {
	resp := &SetLiveWatermarkStatusResponse{}
	err := client.Request("live", "SetLiveWatermarkStatus", "2018-08-01").Do(req, resp)
	return resp, err
}

type SetLiveWatermarkStatusResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

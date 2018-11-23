package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 更新水印
// https://cloud.tencent.com/document/api/267/30150

type UpdateLiveWatermarkRequest struct {
	// 水印图片url。
	PictureUrl string `name:"PictureUrl"`
	// 区域
	Region string `name:"Region"`
	// 水印ID。
	WatermarkId int64 `name:"WatermarkId"`
	// 水印名称。
	WatermarkName *string `name:"WatermarkName,omitempty"`
	// 显示位置，X轴偏移。
	XPosition int64 `name:"XPosition"`
	// 显示位置，Y轴偏移。
	YPosition int64 `name:"YPosition"`
}

func (req *UpdateLiveWatermarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*UpdateLiveWatermarkResponse, error) {
	resp := &UpdateLiveWatermarkResponse{}
	err := client.Request("live", "UpdateLiveWatermark", "2018-08-01").Do(req, resp)
	return resp, err
}

type UpdateLiveWatermarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

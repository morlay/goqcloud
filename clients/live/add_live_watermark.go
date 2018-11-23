package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 添加水印
// https://cloud.tencent.com/document/api/267/30154

type AddLiveWatermarkRequest struct {
	// 水印图片url。
	PictureUrl string `name:"PictureUrl"`
	// 区域
	Region string `name:"Region"`
	// 水印名称。
	WatermarkName string `name:"WatermarkName"`
	// 显示位置,X轴偏移。
	XPosition *int64 `name:"XPosition,omitempty"`
	// 显示位置,Y轴偏移。
	YPosition *int64 `name:"YPosition,omitempty"`
}

func (req *AddLiveWatermarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*AddLiveWatermarkResponse, error) {
	resp := &AddLiveWatermarkResponse{}
	err := client.Request("live", "AddLiveWatermark", "2018-08-01").Do(req, resp)
	return resp, err
}

type AddLiveWatermarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 水印ID。
	WatermarkId int64 `json:"WatermarkId"`
}

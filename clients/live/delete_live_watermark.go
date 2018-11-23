package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除水印
// https://cloud.tencent.com/document/api/267/30153

type DeleteLiveWatermarkRequest struct {
	// 区域
	Region string `name:"Region"`
	// 水印ID。
	WatermarkId int64 `name:"WatermarkId"`
}

func (req *DeleteLiveWatermarkRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeleteLiveWatermarkResponse, error) {
	resp := &DeleteLiveWatermarkResponse{}
	err := client.Request("live", "DeleteLiveWatermark", "2018-08-01").Do(req, resp)
	return resp, err
}

type DeleteLiveWatermarkResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取收银台前顾客身份ID
// https://cloud.tencent.com/document/api/860/19900

type DescribeCameraPersonRequest struct {
	// 摄像头id
	CameraId int64 `name:"CameraId"`
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId string `name:"CompanyId"`
	// 拉取结束时间戳，单位秒，不超过StartTime+10秒，超过默认为StartTime+10
	EndTime int64 `name:"EndTime"`
	// 是否需要base64的图片，0-不需要，1-需要，默认0
	IsNeedPic *int64 `name:"IsNeedPic,omitempty"`
	// 拉取图片数，默认为1，最大为3
	Num *int64 `name:"Num,omitempty"`
	// pos机id
	PosId *string `name:"PosId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId int64 `name:"ShopId"`
	// 拉取开始时间戳，单位秒
	StartTime int64 `name:"StartTime"`
}

func (req *DescribeCameraPersonRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeCameraPersonResponse, error) {
	resp := &DescribeCameraPersonResponse{}
	err := client.Request("youmall", "DescribeCameraPerson", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeCameraPersonResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 摄像机id
	CameraId int64 `json:"CameraId"`
	// 集团id
	CompanyId string `json:"CompanyId"`
	// 抓取的顾客信息
	Infos []*CameraPersonInfo `json:"Infos"`
	// pos机id
	PosId string `json:"PosId"`
	// 店铺id
	ShopId int64 `json:"ShopId"`
}

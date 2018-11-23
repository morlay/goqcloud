package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取已绑定收银台顾客FaceID
// https://cloud.tencent.com/document/api/860/19899

type DescribeFaceIDByTempIDRequest struct {
	// 摄像头id
	CameraId int64 `name:"CameraId"`
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId string `name:"CompanyId"`
	// 图片url过期时间：在当前时间+PictureExpires秒后，图片url无法继续正常访问；单位s；默认值12460*60（1天）
	PictureExpires *int64 `name:"PictureExpires,omitempty"`
	// pos机id
	PosId *string `name:"PosId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取
	ShopId int64 `name:"ShopId"`
	// 临时id
	TempId string `name:"TempId"`
}

func (req *DescribeFaceIDByTempIDRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeFaceIDByTempIDResponse, error) {
	resp := &DescribeFaceIDByTempIDResponse{}
	err := client.Request("youmall", "DescribeFaceIdByTempId", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeFaceIDByTempIDResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 摄像机id
	CameraId int64 `json:"CameraId"`
	// 集团id
	CompanyId string `json:"CompanyId"`
	// 临时id对应的face id
	FaceId int64 `json:"FaceId"`
	// 顾客属性信息
	PersonInfo PersonInfo `json:"PersonInfo"`
	// pos机id
	PosId string `json:"PosId"`
	// 店铺id
	ShopId int64 `json:"ShopId"`
	// 请求的临时id
	TempId string `json:"TempId"`
}

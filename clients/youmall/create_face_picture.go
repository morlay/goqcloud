package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 上传人脸图片
// https://cloud.tencent.com/document/api/860/20420

type CreateFacePictureRequest struct {
	// 集团ID
	CompanyId string `name:"CompanyId"`
	// 是否强制更新：为ture时会为用户创建一个新的指定PersonType的身份;目前这个参数已废弃，可不传
	IsForceUpload *bool `name:"IsForceUpload,omitempty"`
	// 人物类型（0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType int64 `name:"PersonType"`
	// 图片BASE编码
	Picture string `name:"Picture"`
	// 图片名称
	PictureName string `name:"PictureName"`
	// 区域
	Region string `name:"Region"`
	// 店铺ID
	ShopId int64 `name:"ShopId"`
}

func (req *CreateFacePictureRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateFacePictureResponse, error) {
	resp := &CreateFacePictureResponse{}
	err := client.Request("youmall", "CreateFacePicture", "2018-02-28").Do(req, resp)
	return resp, err
}

type CreateFacePictureResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 人物ID
	PersonId int64 `json:"PersonId"`
	// 图片url
	PictureUrl string `json:"PictureUrl"`
	// 0.正常建档 1.重复身份 2.未检测到人脸 3.检测到多个人脸 4.人脸大小过小 5.人脸质量不达标 6.其他错误
	Status int64 `json:"Status"`
}

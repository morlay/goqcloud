package iotcloud

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建设备
// https://cloud.tencent.com/document/api/634/19496

type CreateDeviceRequest struct {
	// 设备属性
	Attribute *Attribute `name:"Attribute,omitempty"`
	// 是否使用自定义PSK，默认不使用
	DefinedPsk *string `name:"DefinedPsk,omitempty"`
	// 设备名称。命名规则：[a-zA-Z0-9:_-]{1,48}。
	DeviceName string `name:"DeviceName"`
	// IMEI，当产品是NB-IoT产品时，此字段必填
	Imei *string `name:"Imei,omitempty"`
	// 运营商类型，当产品是NB-IoT产品时，此字段必填。1表示中国电信，2表示中国移动，3表示中国联通
	Isp *int64 `name:"Isp,omitempty"`
	// LoRa设备的DevEui，当创建LoRa时，此字段必填
	LoraDevEui *string `name:"LoraDevEui,omitempty"`
	// 产品 ID 。创建产品时腾讯云为用户分配全局唯一的 ID
	ProductId string `name:"ProductId"`
	// 区域
	Region string `name:"Region"`
}

func (req *CreateDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDeviceResponse, error) {
	resp := &CreateDeviceResponse{}
	err := client.Request("iotcloud", "CreateDevice", "2018-06-14").Do(req, resp)
	return resp, err
}

type CreateDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 设备证书，用于 TLS 建立链接时校验客户端身份。采用非对称加密时返回该参数
	DeviceCert string `json:"DeviceCert"`
	// 设备名称
	DeviceName string `json:"DeviceName"`
	// 设备私钥，用于 TLS 建立链接时校验客户端身份，腾讯云后台不保存，请妥善保管。采用非对称加密时返回该参数
	DevicePrivateKey string `json:"DevicePrivateKey"`
	// 对称加密密钥，base64编码。采用对称加密时返回该参数
	DevicePsk string `json:"DevicePsk"`
	// LoRa设备的DevEui，当设备是LoRa设备时，会返回该字段
	LoraDevEui string `json:"LoraDevEui"`
}

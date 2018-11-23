package ds

// 签署关键字信息

type SignKeyword struct {
	// 签章图片高度
	ImageHeight string `json:"ImageHeight"`
	// 签章突破宽度
	ImageWidth string `json:"ImageWidth"`
	// 关键字
	Keyword string `json:"Keyword"`
	// X轴偏移坐标
	OffsetCoordX string `json:"OffsetCoordX"`
	// Y轴偏移坐标
	OffsetCoordY string `json:"OffsetCoordY"`
}

// 签署坐标对象

type SignLocation struct {
	// 签名域左下角X轴坐标轴
	SignLocationLBX string `json:"SignLocationLBX"`
	// 签名域左下角Y轴坐标轴
	SignLocationLBY string `json:"SignLocationLBY"`
	// 签名域右上角X轴坐标轴
	SignLocationRUX string `json:"SignLocationRUX"`
	// 签名域右上角Y轴坐标轴
	SignLocationRUY string `json:"SignLocationRUY"`
	// 签名域页数
	SignOnPage string `json:"SignOnPage"`
}

// 签署人信息

type SignInfo struct {
	// 账户ID
	AccountResId string `json:"AccountResId"`
	// 授权时间，格式为年月日时分秒，例20160801095509
	AuthorizationTime string `json:"AuthorizationTime"`
	// 默认值：1  表示RSA证书， 2 表示国密证书， 参数不传时默认为1
	CertType *int64 `json:"CertType,omitempty"`
	// 签名图片，优先级比SealId高
	ImageData *string `json:"ImageData,omitempty"`
	// 授权IP地址
	Location string `json:"Location"`
	// 签章ID
	SealId *string `json:"SealId,omitempty"`
}

package tmt

// 翻译结果

type ItemValue struct {
	// 高度
	H int64 `json:"H"`
	// 识别出的源文
	SourceText string `json:"SourceText"`
	// 翻译后的译文
	TargetText string `json:"TargetText"`
	// 宽度
	W int64 `json:"W"`
	// X坐标
	X int64 `json:"X"`
	// Y坐标
	Y int64 `json:"Y"`
}

// 图片翻译结果

type ImageRecord struct {
	// 图片翻译结果
	Value []*ItemValue `json:"Value"`
}

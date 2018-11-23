package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 门店到访通知回调地址注册
// https://cloud.tencent.com/document/api/860/18456

type RegisterCallbackRequest struct {
	// 通知回调地址，完整url，示例（http://youmall.tencentcloudapi.com/）
	BackUrl string `name:"BackUrl"`
	// 集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId string `name:"CompanyId"`
	// 是否需要顾客图片，1-需要图片，其它-不需要图片
	NeedFacePic *int64 `name:"NeedFacePic,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 请求时间戳
	Time int64 `name:"Time"`
}

func (req *RegisterCallbackRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RegisterCallbackResponse, error) {
	resp := &RegisterCallbackResponse{}
	err := client.Request("youmall", "RegisterCallback", "2018-02-28").Do(req, resp)
	return resp, err
}

type RegisterCallbackResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

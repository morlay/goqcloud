package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 标记顾客身份类型
// https://cloud.tencent.com/document/api/860/19898

type ModifyPersonTagInfoRequest struct {
	// 优mall集团id，通过"指定身份标识获取客户门店列表"接口获取
	CompanyId string `name:"CompanyId"`
	// 区域
	Region string `name:"Region"`
	// 优mall店铺id，通过"指定身份标识获取客户门店列表"接口获取，为0则拉取集团全部店铺当前
	ShopId int64 `name:"ShopId"`
	// 需要设置的顾客信息，批量设置最大为10个
	Tags []*PersonTagInfo `name:"Tags"`
}

func (req *ModifyPersonTagInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyPersonTagInfoResponse, error) {
	resp := &ModifyPersonTagInfoResponse{}
	err := client.Request("youmall", "ModifyPersonTagInfo", "2018-02-28").Do(req, resp)
	return resp, err
}

type ModifyPersonTagInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

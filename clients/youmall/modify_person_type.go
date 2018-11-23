package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改顾客身份类型接口
// https://cloud.tencent.com/document/api/860/30428

type ModifyPersonTypeRequest struct {
	// 集团ID
	CompanyId string `name:"CompanyId"`
	// 顾客ID
	PersonId int64 `name:"PersonId"`
	// 身份子类型:PersonType=0时(普通顾客)，0普通顾客PersonType=1时(白名单)，0店员，1商场人员，2其他类型人员，3区域经理，4注册用户，5VIP用户PersonType=2时(黑名单)，0普通黑名单，1小偷)
	PersonSubType int64 `name:"PersonSubType"`
	// 身份类型(0表示普通顾客，1 白名单，2 表示黑名单）
	PersonType int64 `name:"PersonType"`
	// 区域
	Region string `name:"Region"`
	// 门店ID
	ShopId int64 `name:"ShopId"`
}

func (req *ModifyPersonTypeRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyPersonTypeResponse, error) {
	resp := &ModifyPersonTypeResponse{}
	err := client.Request("youmall", "ModifyPersonType", "2018-02-28").Do(req, resp)
	return resp, err
}

type ModifyPersonTypeResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

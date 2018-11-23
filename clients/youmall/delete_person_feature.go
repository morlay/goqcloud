package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 删除顾客特征
// https://cloud.tencent.com/document/api/860/30429

type DeletePersonFeatureRequest struct {
	// 公司ID
	CompanyId string `name:"CompanyId"`
	// 顾客ID
	PersonId int64 `name:"PersonId"`
	// 区域
	Region string `name:"Region"`
	// 门店ID
	ShopId int64 `name:"ShopId"`
}

func (req *DeletePersonFeatureRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DeletePersonFeatureResponse, error) {
	resp := &DeletePersonFeatureResponse{}
	err := client.Request("youmall", "DeletePersonFeature", "2018-02-28").Do(req, resp)
	return resp, err
}

type DeletePersonFeatureResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询拉流配置
// https://cloud.tencent.com/document/api/267/30158

type DescribePullStreamConfigsRequest struct {
	// 配置id。
	ConfigId *string `name:"ConfigId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribePullStreamConfigsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribePullStreamConfigsResponse, error) {
	resp := &DescribePullStreamConfigsResponse{}
	err := client.Request("live", "DescribePullStreamConfigs", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribePullStreamConfigsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 拉流配置。
	PullStreamConfigs []*PullStreamConfig `json:"PullStreamConfigs"`
}

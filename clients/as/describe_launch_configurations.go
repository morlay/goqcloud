package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询启动配置
// https://cloud.tencent.com/document/api/377/20445

type DescribeLaunchConfigurationsRequest struct {
	// 过滤条件。 launch-configuration-id - String - 是否必填：否 -（过滤条件）按照启动配置ID过滤。 launch-configuration-name - String - 是否必填：否 -（过滤条件）按照启动配置名称过滤。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定LaunchConfigurationIds和Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 按照一个或者多个启动配置ID查询。启动配置ID形如：asc-ouy1ax38。每次请求的上限为100。参数不支持同时指定LaunchConfigurationIds和Filters。
	LaunchConfigurationIds []*string `name:"LaunchConfigurationIds,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeLaunchConfigurationsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLaunchConfigurationsResponse, error) {
	resp := &DescribeLaunchConfigurationsResponse{}
	err := client.Request("as", "DescribeLaunchConfigurations", "2018-04-19").Do(req, resp)
	return resp, err
}

type DescribeLaunchConfigurationsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 启动配置详细信息列表。
	LaunchConfigurationSet []*LaunchConfiguration `json:"LaunchConfigurationSet"`
	// 符合条件的启动配置数量。
	TotalCount int64 `json:"TotalCount"`
}

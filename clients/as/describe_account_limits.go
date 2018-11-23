package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询用户账号的资源限制
// https://cloud.tencent.com/document/api/377/20443

type DescribeAccountLimitsRequest struct {
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeAccountLimitsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeAccountLimitsResponse, error) {
	resp := &DescribeAccountLimitsResponse{}
	err := client.Request("as", "DescribeAccountLimits", "2018-04-19").Do(req, resp)
	return resp, err
}

type DescribeAccountLimitsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 用户账户被允许创建的伸缩组最大数量
	MaxNumberOfAutoScalingGroups int64 `json:"MaxNumberOfAutoScalingGroups"`
	// 用户账户被允许创建的启动配置最大数量
	MaxNumberOfLaunchConfigurations int64 `json:"MaxNumberOfLaunchConfigurations"`
	// 用户账户伸缩组的当前数量
	NumberOfAutoScalingGroups int64 `json:"NumberOfAutoScalingGroups"`
	// 用户账户启动配置的当前数量
	NumberOfLaunchConfigurations int64 `json:"NumberOfLaunchConfigurations"`
}

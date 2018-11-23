package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建伸缩组
// https://cloud.tencent.com/document/api/377/20440

type CreateAutoScalingGroupRequest struct {
	// 伸缩组名称，在您账号中必须唯一。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超55个字节。
	AutoScalingGroupName string `name:"AutoScalingGroupName"`
	// 默认冷却时间，单位秒，默认值为300
	DefaultCooldown *int64 `name:"DefaultCooldown,omitempty"`
	// 期望实例数，大小介于最小实例数和最大实例数之间
	DesiredCapacity *int64 `name:"DesiredCapacity,omitempty"`
	// 应用型负载均衡器列表，目前长度上限为1，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	ForwardLoadBalancers []*ForwardLoadBalancer `name:"ForwardLoadBalancers,omitempty"`
	// 启动配置ID
	LaunchConfigurationId string `name:"LaunchConfigurationId"`
	// 传统负载均衡器ID列表，目前长度上限为1，LoadBalancerIds 和 ForwardLoadBalancers 二者同时最多只能指定一个
	LoadBalancerIds []*string `name:"LoadBalancerIds,omitempty"`
	// 最大实例数，取值范围为0-2000。
	MaxSize int64 `name:"MaxSize"`
	// 最小实例数，取值范围为0-2000。
	MinSize int64 `name:"MinSize"`
	// 项目ID
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 重试策略，取值包括 IMMEDIATE_RETRY 和 INCREMENTAL_INTERVALS，默认取值为 IMMEDIATE_RETRY。 IMMEDIATE_RETRY，立即重试，在较短时间内快速重试，连续失败超过一定次数（5次）后不再重试。 INCREMENTAL_INTERVALS，间隔递增重试，随着连续失败次数的增加，重试间隔逐渐增大，重试间隔从秒级到1天不等。在连续失败超过一定次数（25次）后不再重试。
	RetryPolicy *string `name:"RetryPolicy,omitempty"`
	// 子网ID列表，VPC场景下必须指定子网
	SubnetIds []*string `name:"SubnetIds,omitempty"`
	// 销毁策略，目前长度上限为1
	TerminationPolicies []*string `name:"TerminationPolicies,omitempty"`
	// VPC ID，基础网络则填空字符串
	VpcId string `name:"VpcId"`
	// 可用区列表，基础网络场景下必须指定可用区
	Zones []*string `name:"Zones,omitempty"`
}

func (req *CreateAutoScalingGroupRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateAutoScalingGroupResponse, error) {
	resp := &CreateAutoScalingGroupResponse{}
	err := client.Request("as", "CreateAutoScalingGroup", "2018-04-19").Do(req, resp)
	return resp, err
}

type CreateAutoScalingGroupResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 伸缩组ID
	AutoScalingGroupId string `json:"AutoScalingGroupId"`
}

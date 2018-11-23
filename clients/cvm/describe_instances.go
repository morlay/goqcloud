package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看实例列表
// https://cloud.tencent.com/document/api/213/15728

type DescribeInstancesRequest struct {
	// 过滤条件。 zone - String - 是否必填：否 -（过滤条件）按照可用区过滤。 project-id - Integer - 是否必填：否 -（过滤条件）按照项目ID过滤。可通过调用DescribeProject查询已创建的项目列表或登录控制台进行查看；也可以调用AddProject创建新的项目。 host-id - String - 是否必填：否 - （过滤条件）按照CDH ID过滤。CDH ID形如：host-xxxxxxxx。 vpc-id - String - 是否必填：否 - （过滤条件）按照VPC ID进行过滤。VPC ID形如：vpc-xxxxxxxx。 instance-id - String - 是否必填：否 - （过滤条件）按照实例ID过滤。实例ID形如：ins-xxxxxxxx。 security-group-id - String - 是否必填：否 - （过滤条件）按照安全组ID过滤，安全组ID形如: sg-8jlk3f3r。 instance-name - String - 是否必填：否 - （过滤条件）按照实例名称过滤。 instance-charge-type - String - 是否必填：否 -（过滤条件）按照实例计费模式过滤。 (PREPAID：表示预付费，即包年包月 | POSTPAID_BY_HOUR：表示后付费，即按量计费 | CDHPAID：表示CDH付费，即只对CDH计费，不对CDH上的实例计费。 )   private-ip-address - String - 是否必填：否 - （过滤条件）按照实例主网卡的内网IP过滤。 public-ip-address - String - 是否必填：否 - （过滤条件）按照实例主网卡的公网IP过滤，包含实例创建时自动分配的IP和实例创建后手动绑定的弹性IP。 tag-key - String - 是否必填：否 - （过滤条件）按照标签键进行过滤。 tag-value - String - 是否必填：否 - （过滤条件）按照标签值进行过滤。 tag:tag-key - String - 是否必填：否 - （过滤条件）按照标签键值对进行过滤。 tag-key使用具体的标签键进行替换。使用请参考示例2。每次请求的Filters的上限为10，Filter.Values的上限为5。参数不支持同时指定InstanceIds和Filters。
	Filters []*Filter `name:"Filters,omitempty"`
	// 按照一个或者多个实例ID查询。实例ID形如：ins-xxxxxxxx。（此参数的具体格式可参考API简介的id.N一节）。每次请求的实例的上限为100。参数不支持同时指定InstanceIds和Filters。
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 返回数量，默认为20，最大值为100。关于Limit的更进一步介绍请参考 API 简介中的相关小节。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset的更进一步介绍请参考 API 简介中的相关小节。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeInstancesResponse, error) {
	resp := &DescribeInstancesResponse{}
	err := client.Request("cvm", "DescribeInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息列表。
	InstanceSet []*Instance `json:"InstanceSet"`
	// 符合条件的实例数量。
	TotalCount int64 `json:"TotalCount"`
}

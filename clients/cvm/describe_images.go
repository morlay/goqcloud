package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查看镜像列表
// https://cloud.tencent.com/document/api/213/15715
type DescribeImagesRequest struct {
	// 过滤条件，每次请求的Filters的上限为0，Filters.Values的上限为5。参数不可以同时指定ImageIds和Filters。详细的过滤条件如下： image-id - String - 是否必填： 否 - （过滤条件）按照镜像ID进行过滤 image-type - String - 是否必填： 否 - （过滤条件）按照镜像类型进行过滤。取值范围：详见镜像类型。 image-state - String - 是否必填： 否 - （过滤条件）按照镜像状态进行过滤。取值范围：详见镜像状态。
	Filters []*Filter `name:"Filters,omitempty"`
	// 镜像ID列表 。镜像ID如：img-gvbnzy6f。array型参数的格式可以参考API简介。镜像ID可以通过如下方式获取：通过DescribeImages接口返回的ImageId获取。通过镜像控制台获取。
	ImageIds []*string `name:"ImageIds,omitempty"`
	// 实例类型，如 S1.SMALL1
	InstanceType *string `name:"InstanceType,omitempty"`
	// 数量限制，默认为20，最大值为100。关于Limit详见API简介。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为0。关于Offset详见API简介。
	Offset *int64 `name:"Offset,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeImagesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeImagesResponse, error) {
	resp := &DescribeImagesResponse{}
	err := client.Request("cvm", "DescribeImages", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeImagesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 一个关于镜像详细信息的结构体，主要包括镜像的主要状态与属性。
	ImageSet []*Image `json:"ImageSet"`
	// 符合要求的镜像数量。
	TotalCount int64 `json:"TotalCount"`
}

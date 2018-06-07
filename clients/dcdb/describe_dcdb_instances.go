package dcdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询实例列表
// https://cloud.tencent.com/document/api/557/16140

type DescribeDcdbInstancesRequest struct {
	// 按照一个或者多个实例 ID 查询。实例 ID 形如：dcdbt-2t4cf98d
	InstanceIds []*string `name:"InstanceIds,omitempty"`
	// 是否根据 VPC 网络来搜索，0 为否，1 为是
	IsFilterVpc []*int64 `name:"IsFilterVpc,omitempty"`
	// 返回数量，默认为 10，最大值为 100。
	Limit *int64 `name:"Limit,omitempty"`
	// 偏移量，默认为 0
	Offset *int64 `name:"Offset,omitempty"`
	// 排序字段， projectId， createtime， instancename 三者之一
	OrderBy *string `name:"OrderBy,omitempty"`
	// 排序类型， desc 或者 asc
	OrderByType *string `name:"OrderByType,omitempty"`
	// 按项目 ID 查询
	ProjectIds []*int64 `name:"ProjectIds,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 搜索的关键字，支持模糊搜索。多个关键字使用换行符（'\n'）分割。
	SearchKey *string `name:"SearchKey,omitempty"`
	// 搜索的字段名，当前支持的值有：instancename、vip、all。传 instancename 表示按实例名进行搜索；传 vip 表示按内网IP进行搜索；传 all 将会按实例ID、实例名和内网IP进行搜索。
	SearchName *string `name:"SearchName,omitempty"`
	// 私有网络的子网 ID， IsFilterVpc 为 1 时有效
	SubnetId *string `name:"SubnetId,omitempty"`
	// 私有网络 ID， IsFilterVpc 为 1 时有效
	VpcId *string `name:"VpcId,omitempty"`
}

func (req *DescribeDcdbInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeDcdbInstancesResponse, error) {
	resp := &DescribeDcdbInstancesResponse{}
	err := client.Request("dcdb", "DescribeDCDBInstances", "2018-04-11").Do(req, resp)
	return resp, err
}

type DescribeDcdbInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 实例详细信息列表
	Instances []*DCDBInstanceInfo `json:"Instances"`
	// 符合条件的实例数量
	TotalCount int64 `json:"TotalCount"`
}

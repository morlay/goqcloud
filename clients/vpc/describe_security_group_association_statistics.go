package vpc

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询安全组关联实例统计
// https://cloud.tencent.com/document/api/215/17799

type DescribeSecurityGroupAssociationStatisticsRequest struct {
	// 区域
	Region string `name:"Region"`
	// 安全实例ID，例如sg-33ocnj9n，可通过DescribeSecurityGroups获取。
	SecurityGroupIds []*string `name:"SecurityGroupIds"`
}

func (req *DescribeSecurityGroupAssociationStatisticsRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeSecurityGroupAssociationStatisticsResponse, error) {
	resp := &DescribeSecurityGroupAssociationStatisticsResponse{}
	err := client.Request("vpc", "DescribeSecurityGroupAssociationStatistics", "2017-03-12").Do(req, resp)
	return resp, err
}

type DescribeSecurityGroupAssociationStatisticsResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 安全组关联实例统计。
	SecurityGroupAssociationStatisticsSet []*SecurityGroupAssociationStatistics `json:"SecurityGroupAssociationStatisticsSet"`
}

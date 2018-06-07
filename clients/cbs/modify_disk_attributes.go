package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 修改云硬盘属性
// https://cloud.tencent.com/document/api/362/15659

type ModifyDiskAttributesRequest struct {
	// 一个或多个待操作的云硬盘ID。如果传入多个云盘ID，仅支持所有云盘修改为同一属性。
	DiskIds []*string `name:"DiskIds"`
	// 新的云硬盘名称。
	DiskName *string `name:"DiskName,omitempty"`
	// 是否为弹性云盘，FALSE表示非弹性云盘，TRUE表示弹性云盘。仅支持非弹性云盘修改为弹性云盘。
	Portable *bool `name:"Portable,omitempty"`
	// 新的云硬盘项目ID，只支持修改弹性云盘的项目ID。通过DescribeProject接口查询可用项目及其ID。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *ModifyDiskAttributesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*ModifyDiskAttributesResponse, error) {
	resp := &ModifyDiskAttributesResponse{}
	err := client.Request("cbs", "ModifyDiskAttributes", "2017-03-12").Do(req, resp)
	return resp, err
}

type ModifyDiskAttributesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

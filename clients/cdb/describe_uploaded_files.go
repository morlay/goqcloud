package cdb

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询导入SQL文件列表
// https://cloud.tencent.com/document/api/236/30161

type DescribeUploadedFilesRequest struct {
	// 单次请求返回的数量，默认值为20。
	Limit *int64 `name:"Limit,omitempty"`
	// 记录偏移量，默认值为0。
	Offset *int64 `name:"Offset,omitempty"`
	// 文件路径。该字段应填用户主账号的OwnerUin信息。
	Path string `name:"Path"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeUploadedFilesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeUploadedFilesResponse, error) {
	resp := &DescribeUploadedFilesResponse{}
	err := client.Request("cdb", "DescribeUploadedFiles", "2017-03-20").Do(req, resp)
	return resp, err
}

type DescribeUploadedFilesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 返回的SQL文件列表。
	Items []*SqlFileInfo `json:"Items"`
	// 符合查询条件的SQL文件总数。
	TotalCount int64 `json:"TotalCount"`
}

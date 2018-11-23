package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询直播中的流
// https://cloud.tencent.com/document/api/267/20472

type DescribeLiveStreamOnlineListRequest struct {
	// 应用名称。
	AppName *string `name:"AppName,omitempty"`
	// 您的加速域名。
	DomainName string `name:"DomainName"`
	// 取得第几页，默认1。
	PageNum *int64 `name:"PageNum,omitempty"`
	// 每页大小，最大100。 取值：1~100之前的任意整数。默认值：10
	PageSize *int64 `name:"PageSize,omitempty"`
	// 区域
	Region string `name:"Region"`
}

func (req *DescribeLiveStreamOnlineListRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLiveStreamOnlineListResponse, error) {
	resp := &DescribeLiveStreamOnlineListResponse{}
	err := client.Request("live", "DescribeLiveStreamOnlineList", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamOnlineListResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 正在推送流的信息列表
	OnlineInfo []*StreamOnlineInfo `json:"OnlineInfo"`
	// 分页的页码。
	PageNum int64 `json:"PageNum"`
	// 每页显示的条数。
	PageSize int64 `json:"PageSize"`
	// 符合条件的总个数。
	TotalNum int64 `json:"TotalNum"`
	// 总页数。
	TotalPage int64 `json:"TotalPage"`
}

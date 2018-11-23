package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询推过流的流列表
// https://cloud.tencent.com/document/api/267/20471

type DescribeLiveStreamPublishedListRequest struct {
	// 直播流所属应用名称。
	AppName *string `name:"AppName,omitempty"`
	// 您的域名。
	DomainName string `name:"DomainName"`
	// 结束时间。UTC 格式，例如：2016-06-30T19:00:00Z。不超过当前时间。
	EndTime string `name:"EndTime"`
	// 取得第几页。默认值：1
	PageNum *int64 `name:"PageNum,omitempty"`
	// 分页大小。最大值：100。取值范围：1~100 之前的任意整数。默认值：10
	PageSize *int64 `name:"PageSize,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 起始时间。 UTC 格式，例如：2016-06-29T19:00:00Z。和当前时间相隔不超过7天。
	StartTime string `name:"StartTime"`
}

func (req *DescribeLiveStreamPublishedListRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLiveStreamPublishedListResponse, error) {
	resp := &DescribeLiveStreamPublishedListResponse{}
	err := client.Request("live", "DescribeLiveStreamPublishedList", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamPublishedListResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 分页的页码。
	PageNum int64 `json:"PageNum"`
	// 每页大小
	PageSize int64 `json:"PageSize"`
	// 推流记录信息。
	PublishInfo []*StreamName `json:"PublishInfo"`
	// 符合条件的总个数。
	TotalNum int64 `json:"TotalNum"`
	// 总页数。
	TotalPage int64 `json:"TotalPage"`
}

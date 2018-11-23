package live

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 查询在线推流信息列表
// https://cloud.tencent.com/document/api/267/20473

type DescribeLiveStreamOnlineInfoRequest struct {
	// 取得第几页。默认值：1
	PageNum int64 `name:"PageNum"`
	// 分页大小。最大值：100。取值范围：1~100 之前的任意整数。默认值：10
	PageSize int64 `name:"PageSize"`
	// 区域
	Region string `name:"Region"`
	// 0:未开始推流 1:正在推流 2:服务出错 3:已关闭。
	Status *int64 `name:"Status,omitempty"`
	// 流名称。
	StreamName *string `name:"StreamName,omitempty"`
}

func (req *DescribeLiveStreamOnlineInfoRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeLiveStreamOnlineInfoResponse, error) {
	resp := &DescribeLiveStreamOnlineInfoResponse{}
	err := client.Request("live", "DescribeLiveStreamOnlineInfo", "2018-08-01").Do(req, resp)
	return resp, err
}

type DescribeLiveStreamOnlineInfoResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 分页的页码。
	PageNum int64 `json:"PageNum"`
	// 每页大小
	PageSize int64 `json:"PageSize"`
	// 流信息列表
	StreamInfoList []*StreamInfo `json:"StreamInfoList"`
	// 符合条件的总个数。
	TotalNum int64 `json:"TotalNum"`
	// 总页数。
	TotalPage int64 `json:"TotalPage"`
}

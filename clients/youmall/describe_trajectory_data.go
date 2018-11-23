package youmall

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 获取动线轨迹信息
// https://cloud.tencent.com/document/api/860/20419

type DescribeTrajectoryDataRequest struct {
	// 集团ID
	CompanyId string `name:"CompanyId"`
	// 结束日期，格式yyyy-MM-dd
	EndDate string `name:"EndDate"`
	// 顾客性别顾虑，0是男，1是女，其它代表不分性别
	Gender int64 `name:"Gender"`
	// 限制返回数据的最大条数，最大 400（负数代为 400）
	Limit int64 `name:"Limit"`
	// 区域
	Region string `name:"Region"`
	// 店铺ID
	ShopId int64 `name:"ShopId"`
	// 开始日期，格式yyyy-MM-dd
	StartDate string `name:"StartDate"`
}

func (req *DescribeTrajectoryDataRequest) Invoke(client github_com_morlay_goqcloud.Client) (*DescribeTrajectoryDataResponse, error) {
	resp := &DescribeTrajectoryDataResponse{}
	err := client.Request("youmall", "DescribeTrajectoryData", "2018-02-28").Do(req, resp)
	return resp, err
}

type DescribeTrajectoryDataResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 集团ID
	CompanyId string `json:"CompanyId"`
	// 返回动迹的具体信息
	Data []*TrajectorySunData `json:"Data"`
	// 返回动迹中的总人数
	Person int64 `json:"Person"`
	// 店铺ID
	ShopId int64 `json:"ShopId"`
	// 总人数
	TotalPerson int64 `json:"TotalPerson"`
	// 总动迹数目
	TotalTrajectory int64 `json:"TotalTrajectory"`
	// 返回动迹的数目
	Trajectory int64 `json:"Trajectory"`
}

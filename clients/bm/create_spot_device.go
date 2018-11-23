package bm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建黑石竞价实例
// https://cloud.tencent.com/document/api/386/30259

type CreateSpotDeviceRequest struct {
	// 计算单元类型
	ComputeType string `name:"ComputeType"`
	// 购买的计算单元个数
	GoodsNum int64 `name:"GoodsNum"`
	// 操作系统类型ID
	OsTypeId int64 `name:"OsTypeId"`
	// 设置竞价实例密码。可选参数，没有指定会生成随机密码
	Passwd *string `name:"Passwd,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 用户设置的价格。当为SpotWithPriceLimit竞价策略时有效
	SpotPriceLimit *float64 `name:"SpotPriceLimit,omitempty"`
	// 出价策略。可取值为SpotWithPriceLimit和SpotAsPriceGo。SpotWithPriceLimit，用户设置价格上限，需要传SpotPriceLimit参数， 如果市场价高于用户的指定价格，则购买不成功;  SpotAsPriceGo 是随市场价的策略。
	SpotStrategy string `name:"SpotStrategy"`
	// 子网ID
	SubnetId string `name:"SubnetId"`
	// 私有网络ID
	VpcId string `name:"VpcId"`
	// 可用区名称。如ap-guangzhou-bls-1, 通过DescribeRegions获取
	Zone string `name:"Zone"`
}

func (req *CreateSpotDeviceRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateSpotDeviceResponse, error) {
	resp := &CreateSpotDeviceResponse{}
	err := client.Request("bm", "CreateSpotDevice", "2018-04-23").Do(req, resp)
	return resp, err
}

type CreateSpotDeviceResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 任务ID
	FlowId int64 `json:"FlowId"`
	// 创建的服务器ID
	ResourceIds []*string `json:"ResourceIds"`
}

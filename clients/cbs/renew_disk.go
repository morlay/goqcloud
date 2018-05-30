package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 续费云硬盘
// https://cloud.tencent.com/document/api/362/16319
type RenewDiskRequest struct {
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月云盘的续费时长。在云盘与挂载的实例一起续费的场景下，可以指定参数CurInstanceDeadline，此时云盘会按对齐到实例续费后的到期时间来续费。
	DiskChargePrepaid DiskChargePrepaid `name:"DiskChargePrepaid"`
	// 云硬盘ID， 通过DescribeDisks接口查询。
	DiskId string `name:"DiskId"`
	// 区域
	Region string `name:"Region"`
}

func (req *RenewDiskRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RenewDiskResponse, error) {
	resp := &RenewDiskResponse{}
	err := client.Request("cbs", "RenewDisk", "2017-03-12").Do(req, resp)
	return resp, err
}

type RenewDiskResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
}

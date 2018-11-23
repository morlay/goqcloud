package cbs

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建云硬盘
// https://cloud.tencent.com/document/api/362/16312

type CreateDisksRequest struct {
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。
	ClientToken *string `name:"ClientToken,omitempty"`
	// 可选参数，不传该参数则仅执行挂载操作。传入True时，新创建的云盘将设置为随云主机销毁模式，仅对按量计费云硬盘有效。
	DeleteWithInstance *bool `name:"DeleteWithInstance,omitempty"`
	// 预付费模式，即包年包月相关参数设置。通过该参数指定包年包月云盘的购买时长、是否设置自动续费等属性。创建预付费云盘该参数必传，创建按小时后付费云盘无需传该参数。
	DiskChargePrepaid *DiskChargePrepaid `name:"DiskChargePrepaid,omitempty"`
	// 云硬盘计费类型。PREPAID：预付费，即包年包月POSTPAID_BY_HOUR：按小时后付费各类型价格请参考云硬盘价格总览。
	DiskChargeType string `name:"DiskChargeType"`
	// 创建云硬盘数量，不传则默认为1。单次请求最多可创建的云盘数有限制，具体参见云硬盘使用限制。
	DiskCount *int64 `name:"DiskCount,omitempty"`
	// 云盘显示名称。不传则默认为“未命名”。最大长度不能超60个字节。
	DiskName *string `name:"DiskName,omitempty"`
	// 云硬盘大小，单位为GB。如果传入SnapshotId则可不传DiskSize，此时新建云盘的大小为快照大小如果传入SnapshotId同时传入DiskSize，则云盘大小必须大于或等于快照大小云盘大小取值范围参见云硬盘产品分类的说明。
	DiskSize *int64 `name:"DiskSize,omitempty"`
	// 硬盘介质类型。取值范围：CLOUD_BASIC：表示普通云硬盘CLOUD_PREMIUM：表示高性能云硬盘CLOUD_SSD：表示SSD云硬盘。
	DiskType string `name:"DiskType"`
	// 传入该参数用于创建加密云盘，取值固定为ENCRYPT。
	Encrypt *string `name:"Encrypt,omitempty"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目。若不指定项目，将在默认项目下进行创建。
	Placement Placement `name:"Placement"`
	// 区域
	Region string `name:"Region"`
	// 快照ID，如果传入则根据此快照创建云硬盘，快照类型必须为数据盘快照，可通过DescribeSnapshots接口查询快照，见输出参数DiskUsage解释。
	SnapshotId *string `name:"SnapshotId,omitempty"`
	// 云盘绑定的标签。
	Tags []*Tag `name:"Tags,omitempty"`
}

func (req *CreateDisksRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateDisksResponse, error) {
	resp := &CreateDisksResponse{}
	err := client.Request("cbs", "CreateDisks", "2017-03-12").Do(req, resp)
	return resp, err
}

type CreateDisksResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 创建的云硬盘ID列表。
	DiskIdSet []*string `json:"DiskIdSet"`
}

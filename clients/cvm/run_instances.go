package cvm

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建实例
// https://cloud.tencent.com/document/api/213/15730

type RunInstancesRequest struct {
	// 定时任务。通过该参数可以为实例指定定时任务，目前仅支持定时销毁。
	ActionTimer *ActionTimer `name:"ActionTimer,omitempty"`
	// 用于保证请求幂等性的字符串。该字符串由客户生成，需保证不同请求之间唯一，最大值不超过64个ASCII字符。若不指定该参数，则无法保证请求的幂等性。更多详细信息请参阅：如何保证幂等性。
	ClientToken *string `name:"ClientToken,omitempty"`
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘。支持购买的时候指定11块数据盘，其中最多包含1块LOCAL_BASIC数据盘或者LOCAL_SSD数据盘，最多包含10块CLOUD_BASIC数据盘、CLOUD_PREMIUM数据盘或者CLOUD_SSD数据盘。
	DataDisks []*DataDisk `name:"DataDisks,omitempty"`
	// 容灾组id，仅支持指定一个。
	DisasterRecoverGroupIds []*string `name:"DisasterRecoverGroupIds,omitempty"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `name:"EnhancedService,omitempty"`
	// 云服务器的主机名。点号（.）和短横线（-）不能作为 HostName 的首尾字符，不能连续使用。Windows 实例：名字符长度为[2, 15]，允许字母（不限制大小写）、数字和短横线（-）组成，不支持点号（.），不能全是数字。其他类型（Linux 等）实例：字符长度为[2, 60]，允许支持多个点号，点之间为一段，每段允许字母（不限制大小写）、数字和短横线（-）组成。
	HostName *string `name:"HostName,omitempty"`
	// 指定有效的镜像ID，格式形如img-xxx。镜像类型分为四种：公共镜像自定义镜像共享镜像服务市场镜像可通过以下方式获取可用的镜像ID：公共镜像、自定义镜像、共享镜像的镜像ID可通过登录控制台查询；服务镜像市场的镜像ID可通过云市场查询。通过调用接口 DescribeImages ，取返回信息中的ImageId字段。
	ImageId string `name:"ImageId"`
	// 预付费模式，即包年包月相关参数设置。通过该参数可以指定包年包月实例的购买时长、是否设置自动续费等属性。若指定实例的付费模式为预付费则该参数必传。
	InstanceChargePrepaid *InstanceChargePrepaid `name:"InstanceChargePrepaid,omitempty"`
	// 实例计费类型。PREPAID：预付费，即包年包月POSTPAID_BY_HOUR：按小时后付费CDHPAID：独享子机（基于专用宿主机创建，宿主机部分的资源不收费）SPOTPAID：竞价付费默认值：POSTPAID_BY_HOUR。
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 购买实例数量。包年包月实例取值范围：[1，300]，按量计费实例取值范围：[1，100]。默认取值：1。指定购买实例的数量不能超过用户所能购买的剩余配额数量，具体配额相关限制详见CVM实例购买限制。
	InstanceCount *int64 `name:"InstanceCount,omitempty"`
	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `name:"InstanceMarketOptions,omitempty"`
	// 实例显示名称。不指定实例显示名称则默认显示‘未命名’。购买多台实例，如果指定模式串{R:x}，表示生成数字[x, x+n-1]，其中n表示购买实例的数量，例如server_{R:3}，购买1台时，实例显示名称为server_3；购买2台时，实例显示名称分别为server_3，server_4。支持指定多个模式串{R:x}。购买多台实例，如果不指定模式串，则在实例显示名称添加后缀1、2...n，其中n表示购买实例的数量，例如server_，购买2台时，实例显示名称分别为server_1，server_2。
	InstanceName *string `name:"InstanceName,omitempty"`
	// 实例机型。不同实例机型指定了不同的资源规格。对于付费模式为PREPAID或POSTPAID_BY_HOUR的实例创建，具体取值可通过调用接口DescribeInstanceTypeConfigs来获得最新的规格表或参见实例类型描述。若不指定该参数，则默认机型为S1.SMALL1。对于付费模式为CDHPAID的实例创建，该参数以"CDH_"为前缀，根据cpu和内存配置生成，具体形式为：CDH_XCXG，例如对于创建cpu为1核，内存为1G大小的专用宿主机的实例，该参数应该为CDH_1C1G。
	InstanceType *string `name:"InstanceType,omitempty"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `name:"InternetAccessible,omitempty"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `name:"LoginSettings,omitempty"`
	// 实例所在的位置。通过该参数可以指定实例所属可用区，所属项目，所属宿主机（在专用宿主机上创建子机时指定）等属性。
	Placement Placement `name:"Placement"`
	// 区域
	Region string `name:"Region"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的sgId字段来获取。若不指定该参数，则绑定默认安全组。
	SecurityGroupIds []*string `name:"SecurityGroupIds,omitempty"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `name:"SystemDisk,omitempty"`
	// 标签描述列表。通过指定该参数可以同时绑定标签到相应的资源实例，当前仅支持绑定标签到云主机实例。
	TagSpecification []*TagSpecification `name:"TagSpecification,omitempty"`
	// 提供给实例使用的用户数据，需要以 base64 方式编码，支持的最大数据大小为 16KB。关于获取此参数的详细介绍，请参阅Windows和Linux启动时运行命令。
	UserData *string `name:"UserData,omitempty"`
	// 私有网络相关信息配置。通过该参数可以指定私有网络的ID，子网ID等信息。若不指定该参数，则默认使用基础网络。若在此参数中指定了私有网络ip，表示每个实例的主网卡ip，而且InstanceCount参数必须与私有网络ip的个数一致。
	VirtualPrivateCloud *VirtualPrivateCloud `name:"VirtualPrivateCloud,omitempty"`
}

func (req *RunInstancesRequest) Invoke(client github_com_morlay_goqcloud.Client) (*RunInstancesResponse, error) {
	resp := &RunInstancesResponse{}
	err := client.Request("cvm", "RunInstances", "2017-03-12").Do(req, resp)
	return resp, err
}

type RunInstancesResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 当通过本接口来创建实例时会返回该参数，表示一个或多个实例ID。返回实例ID列表并不代表实例创建成功，可根据 DescribeInstances 接口查询返回的InstancesSet中对应实例的ID的状态来判断创建是否完成；如果实例状态由“准备中”变为“正在运行”，则为创建成功。
	InstanceIdSet []*string `json:"InstanceIdSet"`
}

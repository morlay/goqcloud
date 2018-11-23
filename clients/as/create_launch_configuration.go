package as

import (
	github_com_morlay_goqcloud "github.com/morlay/goqcloud"
)

// 创建启动配置
// https://cloud.tencent.com/document/api/377/20447

type CreateLaunchConfigurationRequest struct {
	// 实例数据盘配置信息。若不指定该参数，则默认不购买数据盘，最多支持指定11块数据盘。
	DataDisks []*DataDisk `name:"DataDisks,omitempty"`
	// 增强服务。通过该参数可以指定是否开启云安全、云监控等服务。若不指定该参数，则默认开启云监控、云安全服务。
	EnhancedService *EnhancedService `name:"EnhancedService,omitempty"`
	// 指定有效的镜像ID，格式形如img-8toqc6s3。镜像类型分为四种：公共镜像自定义镜像共享镜像服务市场镜像可通过以下方式获取可用的镜像ID：公共镜像、自定义镜像、共享镜像的镜像ID可通过登录控制台查询；服务镜像市场的镜像ID可通过云市场查询。通过调用接口 DescribeImages ，取返回信息中的ImageId字段。
	ImageId string `name:"ImageId"`
	// 实例计费类型，CVM默认值按照POSTPAID_BY_HOUR处理。POSTPAID_BY_HOUR：按小时后付费SPOTPAID：竞价付费
	InstanceChargeType *string `name:"InstanceChargeType,omitempty"`
	// 实例的市场相关选项，如竞价实例相关参数，若指定实例的付费模式为竞价付费则该参数必传。
	InstanceMarketOptions *InstanceMarketOptionsRequest `name:"InstanceMarketOptions,omitempty"`
	// 实例机型。不同实例机型指定了不同的资源规格，具体取值可通过调用接口 DescribeInstanceTypeConfigs 来获得最新的规格表或参见实例类型描述。InstanceType和InstanceTypes参数互斥，二者必填一个且只能填写一个。
	InstanceType *string `name:"InstanceType,omitempty"`
	// 实例机型列表，不同实例机型指定了不同的资源规格，最多支持5中实例机型。InstanceType和InstanceTypes参数互斥，二者必填一个且只能填写一个。
	InstanceTypes []*string `name:"InstanceTypes,omitempty"`
	// 公网带宽相关信息设置。若不指定该参数，则默认公网带宽为0Mbps。
	InternetAccessible *InternetAccessible `name:"InternetAccessible,omitempty"`
	// 启动配置显示名称。名称仅支持中文、英文、数字、下划线、分隔符"-"、小数点，最大长度不能超60个字节。
	LaunchConfigurationName string `name:"LaunchConfigurationName"`
	// 实例登录设置。通过该参数可以设置实例的登录方式密码、密钥或保持镜像的原始登录设置。默认情况下会随机生成密码，并以站内信方式知会到用户。
	LoginSettings *LoginSettings `name:"LoginSettings,omitempty"`
	// 实例所属项目ID。该参数可以通过调用 DescribeProject 的返回值中的projectId字段来获取。不填为默认项目。
	ProjectId *int64 `name:"ProjectId,omitempty"`
	// 区域
	Region string `name:"Region"`
	// 实例所属安全组。该参数可以通过调用 DescribeSecurityGroups 的返回值中的SecurityGroupId字段来获取。若不指定该参数，则默认不绑定安全组。
	SecurityGroupIds []*string `name:"SecurityGroupIds,omitempty"`
	// 实例系统盘配置信息。若不指定该参数，则按照系统默认值进行分配。
	SystemDisk *SystemDisk `name:"SystemDisk,omitempty"`
	// 经过 Base64 编码后的自定义数据，最大长度不超过16KB。
	UserData *string `name:"UserData,omitempty"`
}

func (req *CreateLaunchConfigurationRequest) Invoke(client github_com_morlay_goqcloud.Client) (*CreateLaunchConfigurationResponse, error) {
	resp := &CreateLaunchConfigurationResponse{}
	err := client.Request("as", "CreateLaunchConfiguration", "2018-04-19").Do(req, resp)
	return resp, err
}

type CreateLaunchConfigurationResponse struct {
	github_com_morlay_goqcloud.TencentCloudBaseResponse
	// 当通过本接口来创建启动配置时会返回该参数，表示启动配置ID。
	LaunchConfigurationId string `json:"LaunchConfigurationId"`
}

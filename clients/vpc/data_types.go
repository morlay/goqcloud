package vpc

import (
	time "time"
)

// IP地址模板集合

type AddressTemplateGroup struct {
	AddressTemplateGroupId   string    `json:"AddressTemplateGroupId"`
	AddressTemplateGroupName string    `json:"AddressTemplateGroupName"`
	AddressTemplateIdSet     []*string `json:"AddressTemplateIdSet"`
	CreatedTime              string    `json:"CreatedTime"`
}

// 私有网络和基础网络互通设备

type ClassicLinkInstance struct {
	InstanceId string `json:"InstanceId"`
	VpcId      string `json:"VpcId"`
}

// 预付费（包年包月）计费对象。

type InstanceChargePrepaid struct {
	// 购买实例的时长，单位：月。取值范围：1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 24, 36。
	Period int64 `json:"Period"`
	// 自动续费标识。取值范围： NOTIFY_AND_AUTO_RENEW：通知过期且自动续费， NOTIFY_AND_MANUAL_RENEW：通知过期不自动续费。默认：NOTIFY_AND_MANUAL_RENEW
	RenewFlag *string `json:"RenewFlag,omitempty"`
}

// 单项计费价格信息

type ItemPrice struct {
	ChargeUnit    string  `json:"ChargeUnit"`
	DiscountPrice float64 `json:"DiscountPrice"`
	OriginalPrice float64 `json:"OriginalPrice"`
	UnitPrice     float64 `json:"UnitPrice"`
}

// SecurityPolicyDatabase策略

type SecurityPolicyDatabase struct {
	// 本端网段
	LocalCidrBlock string `json:"LocalCidrBlock"`
	// 对端网段
	RemoteCidrBlock []*string `json:"RemoteCidrBlock"`
}

// 价格

type Price struct {
	BandwidthPrice ItemPrice `json:"BandwidthPrice"`
	InstancePrice  ItemPrice `json:"InstancePrice"`
}

// 安全组规则对象

type SecurityGroupPolicy struct {
	// ACCEPT 或 DROP。
	Action *string `json:"Action,omitempty"`
	// IP地址ID或者ID地址组ID。
	AddressTemplate *string `json:"AddressTemplate,omitempty"`
	// 网段或IP(互斥)。
	CidrBlock *string `json:"CidrBlock,omitempty"`
	// 安全组规则描述。
	PolicyDescription *string `json:"PolicyDescription,omitempty"`
	// 安全组规则索引号。
	PolicyIndex *int64 `json:"PolicyIndex,omitempty"`
	// 端口(all, 离散port,  range)。
	Port *string `json:"Port,omitempty"`
	// 协议, 取值: TCP,UDP, ICMP。
	Protocol *string `json:"Protocol,omitempty"`
	// 已绑定安全组的网段或IP。
	SecurityGroupId *string `json:"SecurityGroupId,omitempty"`
	// 协议端口ID或者协议端口组ID。ServiceTemplate和Protocol+Port互斥。
	ServiceTemplate []*string `json:"ServiceTemplate,omitempty"`
}

// VPN通道对象。

type VpnConnection struct {
	CreatedTime               time.Time                 `json:"CreatedTime"`
	CustomerGatewayId         string                    `json:"CustomerGatewayId"`
	EncryptProto              string                    `json:"EncryptProto"`
	IKEOptionsSpecification   IKEOptionsSpecification   `json:"IKEOptionsSpecification"`
	IPSECOptionsSpecification IPSECOptionsSpecification `json:"IPSECOptionsSpecification"`
	NetStatus                 string                    `json:"NetStatus"`
	PreShareKey               string                    `json:"PreShareKey"`
	RouteType                 string                    `json:"RouteType"`
	SecurityPolicyDatabaseSet []*SecurityPolicyDatabase `json:"SecurityPolicyDatabaseSet"`
	State                     string                    `json:"State"`
	VpcId                     string                    `json:"VpcId"`
	VpnConnectionId           string                    `json:"VpnConnectionId"`
	VpnConnectionName         string                    `json:"VpnConnectionName"`
	VpnGatewayId              string                    `json:"VpnGatewayId"`
	VpnProto                  string                    `json:"VpnProto"`
}

// IKE配置（Internet Key Exchange，因特网密钥交换），IKE具有一套自保护机制，用户配置网络安全协议

type IKEOptionsSpecification struct {
	// DH group，指定IKE交换密钥时使用的DH组，可选值：'GROUP1', 'GROUP2', 'GROUP5', 'GROUP14', 'GROUP24'，
	DhGroupName *string `json:"DhGroupName,omitempty"`
	// 协商模式：可选值：'AGGRESSIVE', 'MAIN'，默认为MAIN
	ExchangeMode *string `json:"ExchangeMode,omitempty"`
	// IKE SA Lifetime，单位：秒，设置IKE SA的生存周期，取值范围：60-604800
	IKESaLifetimeSeconds *int64 `json:"IKESaLifetimeSeconds,omitempty"`
	// IKE版本
	IKEVersion *string `json:"IKEVersion,omitempty"`
	// 本端标识，当LocalIdentity选为ADDRESS时，LocalAddress必填。localAddress默认为vpn网关公网IP
	LocalAddress *string `json:"LocalAddress,omitempty"`
	// 本端标识，当LocalIdentity选为FQDN时，LocalFqdnName必填
	LocalFqdnName *string `json:"LocalFqdnName,omitempty"`
	// 本端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	LocalIdentity *string `json:"LocalIdentity,omitempty"`
	// 认证算法：可选值：'MD5', 'SHA1'，默认为MD5
	PropoAuthenAlgorithm *string `json:"PropoAuthenAlgorithm,omitempty"`
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBS-192', 'AES-CBC-256', 'DES-CBC'，默认为3DES-CBC
	PropoEncryAlgorithm *string `json:"PropoEncryAlgorithm,omitempty"`
	// 对端标识，当RemoteIdentity选为ADDRESS时，RemoteAddress必填
	RemoteAddress *string `json:"RemoteAddress,omitempty"`
	// 对端标识，当remoteIdentity选为FQDN时，RemoteFqdnName必填
	RemoteFqdnName *string `json:"RemoteFqdnName,omitempty"`
	// 对端标识类型：可选值：'ADDRESS', 'FQDN'，默认为ADDRESS
	RemoteIdentity *string `json:"RemoteIdentity,omitempty"`
}

// 路由表对象

type RouteTable struct {
	AssociationSet []*RouteTableAssociation `json:"AssociationSet"`
	CreatedTime    string                   `json:"CreatedTime"`
	Main           bool                     `json:"Main"`
	RouteSet       []*Route                 `json:"RouteSet"`
	RouteTableId   string                   `json:"RouteTableId"`
	RouteTableName string                   `json:"RouteTableName"`
	VpcId          string                   `json:"VpcId"`
}

// 协议端口模板集合

type ServiceTemplateGroup struct {
	CreatedTime              string    `json:"CreatedTime"`
	ServiceTemplateGroupId   string    `json:"ServiceTemplateGroupId"`
	ServiceTemplateGroupName string    `json:"ServiceTemplateGroupName"`
	ServiceTemplateIdSet     []*string `json:"ServiceTemplateIdSet"`
}

// 过滤器键值对

type FilterObject struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name string `json:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values"`
}

// 描述了配额信息

type Quota struct {
	QuotaCurrent int64  `json:"QuotaCurrent"`
	QuotaId      string `json:"QuotaId"`
	QuotaLimit   int64  `json:"QuotaLimit"`
}

// 路由表关联关系

type RouteTableAssociation struct {
	RouteTableId string `json:"RouteTableId"`
	SubnetId     string `json:"SubnetId"`
}

// 安全组规则集合

type SecurityGroupPolicySet struct {
	// 出站规则。
	Egress []*SecurityGroupPolicy `json:"Egress,omitempty"`
	// 入站规则。
	Ingress []*SecurityGroupPolicy `json:"Ingress,omitempty"`
	// 安全组规则当前版本。用户每次更新安全规则版本会自动加1，防止更新的路由规则已过期，不填不考虑冲突。
	Version *string `json:"Version,omitempty"`
}

// IP地址模板

type AddressTemplate struct {
	AddressSet          []*string `json:"AddressSet"`
	AddressTemplateId   string    `json:"AddressTemplateId"`
	AddressTemplateName string    `json:"AddressTemplateName"`
	CreatedTime         string    `json:"CreatedTime"`
}

// 对端网关厂商信息对象。

type CustomerGatewayVendor struct {
	// 平台。
	Platform string `json:"Platform"`
	// 软件版本。
	SoftwareVersion string `json:"SoftwareVersion"`
	// 供应商名称。
	VendorName string `json:"VendorName"`
}

// 弹性网卡

type NetworkInterface struct {
	Attachment                  []*InstanceChargePrepaid         `json:"Attachment"`
	CreatedTime                 string                           `json:"CreatedTime"`
	GroupSet                    []*string                        `json:"GroupSet"`
	MacAddress                  string                           `json:"MacAddress"`
	NetworkInterfaceDescription string                           `json:"NetworkInterfaceDescription"`
	NetworkInterfaceId          string                           `json:"NetworkInterfaceId"`
	NetworkInterfaceName        string                           `json:"NetworkInterfaceName"`
	Primary                     bool                             `json:"Primary"`
	PrivateIpAddressSet         []*PrivateIpAddressSpecification `json:"PrivateIpAddressSet"`
	State                       string                           `json:"State"`
	SubnetId                    string                           `json:"SubnetId"`
	VpcId                       string                           `json:"VpcId"`
	Zone                        string                           `json:"Zone"`
}

// 安全组对象

type SecurityGroup struct {
	// 安全组创建时间。
	CreatedTime *string `json:"CreatedTime,omitempty"`
	// 是否是默认安全组，默认安全组不支持删除。
	IsDefault *bool `json:"IsDefault,omitempty"`
	// 项目id，默认0。可在qcloud控制台项目管理页面查询到。
	ProjectId *string `json:"ProjectId,omitempty"`
	// 安全组备注，最多100个字符。
	SecurityGroupDesc string `json:"SecurityGroupDesc"`
	// 安全组实例ID，例如：sg-ohuuioma。
	SecurityGroupId string `json:"SecurityGroupId"`
	// 安全组名称，可任意命名，但不得超过60个字符。
	SecurityGroupName string `json:"SecurityGroupName"`
}

// 子网对象

type Subnet struct {
	AvailableIpAddressCount int64  `json:"AvailableIpAddressCount"`
	CidrBlock               string `json:"CidrBlock"`
	CreatedTime             string `json:"CreatedTime"`
	EnableBroadcast         bool   `json:"EnableBroadcast"`
	IsDefault               bool   `json:"IsDefault"`
	RouteTableId            string `json:"RouteTableId"`
	SubnetId                string `json:"SubnetId"`
	SubnetName              string `json:"SubnetName"`
	VpcId                   string `json:"VpcId"`
	Zone                    string `json:"Zone"`
}

// VPN网关对象。

type VpnGateway struct {
	CreatedTime             time.Time `json:"CreatedTime"`
	ExpiredTime             time.Time `json:"ExpiredTime"`
	InstanceChargeType      string    `json:"InstanceChargeType"`
	InternetMaxBandwidthOut int64     `json:"InternetMaxBandwidthOut"`
	IsAddressBlocked        bool      `json:"IsAddressBlocked"`
	NewPurchasePlan         string    `json:"NewPurchasePlan"`
	PublicIpAddress         string    `json:"PublicIpAddress"`
	RenewFlag               string    `json:"RenewFlag"`
	RestrictState           string    `json:"RestrictState"`
	State                   string    `json:"State"`
	Type                    string    `json:"Type"`
	VpcId                   string    `json:"VpcId"`
	VpnGatewayId            string    `json:"VpnGatewayId"`
	VpnGatewayName          string    `json:"VpnGatewayName"`
}

// 描述 EIP 信息

type Address struct {
	AddressId        string    `json:"AddressId"`
	AddressIp        string    `json:"AddressIp"`
	AddressName      string    `json:"AddressName"`
	AddressStatus    string    `json:"AddressStatus"`
	BindedResourceId string    `json:"BindedResourceId"`
	CreatedTime      time.Time `json:"CreatedTime"`
}

// 过滤器

type Filter struct {
	// 属性名称, 若存在多个Filter时，Filter间的关系为逻辑与（AND）关系。
	Name string `json:"Name"`
	// 属性值, 若同一个Filter存在多个Values，同一Filter下Values间的关系为逻辑或（OR）关系。
	Values []*string `json:"Values"`
}

// IPSec配置，腾讯云提供IPSec安全会话设置

type IPSECOptionsSpecification struct {
	// 加密算法，可选值：'3DES-CBC', 'AES-CBC-128', 'AES-CBC-192', 'AES-CBC-256', 'DES-CBC', 'NULL'， 默认为AES-CBC-128
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty"`
	// IPsec SA lifetime(s)：单位秒，取值范围：180-604800
	IPSECSaLifetimeSeconds *int64 `json:"IPSECSaLifetimeSeconds,omitempty"`
	// IPsec SA lifetime(KB)：单位KB，取值范围：2560-604800
	IPSECSaLifetimeTraffic *int64 `json:"IPSECSaLifetimeTraffic,omitempty"`
	// 认证算法：可选值：'MD5', 'SHA1'，默认为
	IntegrityAlgorith *string `json:"IntegrityAlgorith,omitempty"`
	// PFS：可选值：'NULL', 'DH-GROUP1', 'DH-GROUP2', 'DH-GROUP5', 'DH-GROUP14', 'DH-GROUP24'，默认为NULL
	PfsDhGroup *string `json:"PfsDhGroup,omitempty"`
}

// 私有网络(VPC)对象。

type Vpc struct {
	CidrBlock       string `json:"CidrBlock"`
	CreatedTime     string `json:"CreatedTime"`
	EnableMulticast bool   `json:"EnableMulticast"`
	IsDefault       bool   `json:"IsDefault"`
	VpcId           string `json:"VpcId"`
	VpcName         string `json:"VpcName"`
}

// 对端网关

type CustomerGateway struct {
	CreatedTime         string `json:"CreatedTime"`
	CustomerGatewayId   string `json:"CustomerGatewayId"`
	CustomerGatewayName string `json:"CustomerGatewayName"`
	IpAddress           string `json:"IpAddress"`
}

// 内网IP信息

type PrivateIpAddressSpecification struct {
	// EIP实例ID，例如：eip-11112222。
	AddressId *string `json:"AddressId,omitempty"`
	// 内网IP描述信息。
	Description *string `json:"Description,omitempty"`
	// 公网IP是否被封堵。
	IsWanIpBlocked *bool `json:"IsWanIpBlocked,omitempty"`
	// 是否是主IP。
	Primary *bool `json:"Primary,omitempty"`
	// 内网IP地址。
	PrivateIpAddress string `json:"PrivateIpAddress"`
	// 公网IP地址。
	PublicIpAddress *string `json:"PublicIpAddress,omitempty"`
}

// 路由策略对象

type Route struct {
	DestinationCidrBlock string `json:"DestinationCidrBlock"`
	GatewayId            string `json:"GatewayId"`
	GatewayType          string `json:"GatewayType"`
	RouteDescription     string `json:"RouteDescription"`
	RouteId              int64  `json:"RouteId"`
}

// 协议端口模板

type ServiceTemplate struct {
	CreatedTime         string    `json:"CreatedTime"`
	ServiceSet          []*string `json:"ServiceSet"`
	ServiceTemplateId   string    `json:"ServiceTemplateId"`
	ServiceTemplateName string    `json:"ServiceTemplateName"`
}

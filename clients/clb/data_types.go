package clb

// 监听器后端绑定机器的详细信息

type Backend *Backend

// 证书相关信息

type CertificateOutput *CertificateOutput

// 健康检查信息

type HealthCheck *HealthCheck

// 监听器上注册的后端机器的信息

type ListenerBackend *ListenerBackend

// HTTP/HTTPS转发规则（输入）

type RuleInput *RuleInput

// HTTP/HTTPS转发规则（输出）

type RuleOutput *RuleOutput

// HTTP/HTTPS监听器下的转发规则的机器绑定信息

type RuleTargets *RuleTargets

// 证书信息

type CertificateInput *CertificateInput

// 监听器的信息

type Listener *Listener

// 负载均衡实例的信息

type LoadBalancer *LoadBalancer

// 转发目标，即绑定在负载均衡上的后端机器

type Target *Target

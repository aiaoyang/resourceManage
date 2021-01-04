package resource

// Type 资源类型枚举体
type Type int

// Stat 资源状态
type Stat int

// Info 信息
type Info struct {
	Name      string `json:"name"`
	EndOfTime string `json:"end"`
	Type      string `json:"type"`
	Detail    string `json:"detail"`
	Account   string `json:"account"`
	Index     string `json:"index"`
	Status    Stat   `json:"status"`
}

const (
	// Green 0： 一个月以上
	Green Stat = iota

	// Yellow 1： 一个月以内，一周以上
	Yellow

	// Red 2： 一周以内，未过期
	Red

	// NearDead 3： 已过期
	NearDead
)
const (
	// EcsType ecs资源类型
	EcsType Type = iota

	// RdsType rds资源类型
	RdsType

	// DomainType 域名资源类型
	DomainType

	// CertType 证书资源类型
	CertType

	// AlertType 告警资源类型
	AlertType
)

var (
	// ResourceMap 资源类型名称
	ResourceMap = map[int]string{
		0: "ECS",
		1: "RDS",
		2: "Domain",
		3: "Cert",
		4: "Alert",
	}
)

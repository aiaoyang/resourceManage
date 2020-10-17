package apis

import (
	"net/url"

	"github.com/jinzhu/gorm/dialects/postgres"
)

type ResourceStatus int

const (
	Green ResourceStatus = iota
	Yellow
	Red
	Pendding
)

type Type int

const (
	useless Type = iota
	ECS
	RDS
	DOMAIN
)

type Resource interface {
	String() string
	ID() string
	Name() string
	Status() ResourceStatus
	Type() Type
	Depends() []Resource
	EndOfLife() string
	Values() string
}

type MyResource struct {
	ID     int            `gorm:"AUTO_INCREMENT" json:"id"`
	Name   string         `gorm:"index:name" json:"name"`
	Depend string         `gorm:"depend" json:"depend"`
	Type   string         `gorm:"type" json:"type"`
	Value  postgres.Jsonb `gorm:"default:'{}'"`
	Values url.Values     `gorm:"-" json:"values"`
}

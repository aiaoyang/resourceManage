package aliyun

type stat int

const (
	// 0： 一个月以上
	green stat = iota

	// 1： 一个月以内，一周以上
	yellow

	// 2： 一周以内，未过期
	red

	// 3： 已过期
	nearDead
)

package Ticket

//常量定义
const (
	UNPROCESSED = 0
	CONFIRMED   = 1
	DECLINED    = 2
)

type Bill interface {
	GetNextPay() int64
}

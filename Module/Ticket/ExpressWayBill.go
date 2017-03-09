package Ticket

import (
	. "Gyl/Module/System"
	"time"
)

type ExpressWayBillTicket struct {
	TicketNumber    TicketNum
	PurchaseOrderID TicketNum
	TotalAmount     int64
	TotalWeight     int64
	Orderdate       time.Time

	LogisticsID UserID
	Price       int64

	FromAddr string
	ToAddr   string
	FromName string
	toName   string

	LogisticConfirmDate             time.Time
	LogisticGetDate                 time.Time
	LogisticTransportationStartDate time.Time

	Status int
}

func (u *ExpressWayBillTicket) InitEWTicket() ExpressWayBillTicket {
	//TODO:
	return ExpressWayBillTicket{}
}

func (u *ExpressWayBillTicket) GetEWTicketByTicketID() ExpressWayBillTicket {
	//TODO:
	return ExpressWayBillTicket{}
}

func (u *ExpressWayBillTicket) LogisticInformTransportationStart() error {
	//TODO:
	return nil
}

func (u *ExpressWayBillTicket) LogisticTicketConfirmed() error {
	//TODO:
	u.LogisticConfirmDate = time.Now()
	u.Status = CONFIRMED
	return nil
}

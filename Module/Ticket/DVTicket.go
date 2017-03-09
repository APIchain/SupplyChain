package Ticket

import (
	. "Gyl/Module/System"
	"time"
)

type DVTicket struct {
	TicketNumber           TicketNum
	PurchaseOrderID        TicketNum
	TotalAmount            int64
	TotalWeight            int64
	SentOutDate            time.Time
	LogisticsID            UserID
	ExpressWayBillTicketID ExpressWayBillTicket
	SupplyerID             UserID
}

func (u *DVTicket) InitDVTicket() DVTicket {
	//TODO:
	return DVTicket{}
}

func (u *DVTicket) GetDVTicketByTicketID() DVTicket {
	//TODO:
	return DVTicket{}
}

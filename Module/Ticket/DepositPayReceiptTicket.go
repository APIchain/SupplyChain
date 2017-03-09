package Ticket

import (
	. "Gyl/Module/System"
)

type DepositPayReceiptTicket struct {
	TicketNumber TicketNum
	OrderID      TicketNum
	Paydate      string
	Amount       int64
}

func (u *DepositPayReceiptTicket) InitDepositPayReceiptTicket() AccountsReceivableTicket {
	//TODO:
	return AccountsReceivableTicket{}
}

func (u *DepositPayReceiptTicket) GetDepositPayReceiptTicketByTicketID() DepositPayReceiptTicket {
	//TODO:
	return DepositPayReceiptTicket{}
}

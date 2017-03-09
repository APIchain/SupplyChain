package Ticket

import (
	. "Gyl/Module/System"
)

type AccountsReceivableTicket struct {
	TicketNumber    TicketNum
	PurchaseOrderID TicketNum
	PayEnddate      string
	Amount          int64
}

func (u *AccountsReceivableTicket) InitAccountsReceivableTicket() AccountsReceivableTicket {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return AccountsReceivableTicket{}
}

func (u *AccountsReceivableTicket) GetAccountsReceivableTicketByTicketID() AccountsReceivableTicket {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return AccountsReceivableTicket{}
}

package Ticket

import (
	. "Gyl/Module/System"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

type WVTicket struct {
	TicketNumber    TicketNum
	PurchaseOrderID TicketNum
	DVTicketID      TicketNum
	//Products ProductPurchaseDetail  todo later
	ArriveDate       time.Time
	TotalAmount      int64
	TotalWeight      int64
	SupplyID         UserID
	BuyerConfirmDate time.Time
	Status           int
}

func InitWVTicket(stub shim.ChaincodeStubInterface, purchaseOrderID TicketNum, dvTicket TicketNum, totalAmount int64, totalWeight int64) (*WVTicket, error) {
	//TODO:
	var err error
	WV := new(WVTicket)
	WV.TicketNumber, err = GenTicketNum(stub)
	if err != nil {
		return nil, err
	}
	WV.PurchaseOrderID = purchaseOrderID
	WV.DVTicketID = dvTicket
	WV.TotalAmount = totalAmount
	WV.TotalWeight = totalWeight

	return WV, err
}

func GetWVTicketByTicketID(TicketNum TicketNum) (*WVTicket, error) {
	//TODO:
	return &WVTicket{}, nil
}

func (u *WVTicket) BuyerConfirm() {
	//TODO:
	u.BuyerConfirmDate = time.Now()
	u.Status = CONFIRMED
}

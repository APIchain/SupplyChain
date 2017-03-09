package Control

import (
	. "Gyl/Module/System"
	"Gyl/Module/Roll"
	"Gyl/Module/Ticket"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"Gyl/Module/List"
)

type Bank struct {
	Roll.Bank
}

//支付账单
func (b *Bank) GetNeedConfirmPayTicket(stub shim.ChaincodeStubInterface) (*[]List.ConfirmDetail, error) {
	//TODO:
	return nil, nil
}

func (b *Bank) ConfirmPay(stub shim.ChaincodeStubInterface, tickid TicketNum) error {
	//TODO:
	return nil
}

func (b *Bank) PayReceiptGen() *Ticket.DepositPayReceiptTicket {
	//TODO:
	return &Ticket.DepositPayReceiptTicket{}
}

func (b *Bank) InformPayArrivedWithReceiptionID() error {
	//TODO:
	return nil
}

func GetAllBank(stub shim.ChaincodeStubInterface) *[]Bank {
	//TODO:
	return nil
}

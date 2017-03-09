package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Buyer struct {
	UserID         UserID      `json:"UserID"`
	PurchaseOrders []TicketNum `json:"PurchaseOrders"`
}

func NewBuyerWithInit(stub shim.ChaincodeStubInterface, userID *UserID) (*Buyer, error) {
	buyer := new(Buyer)
	buyer.UserID = *userID
	buyer.PurchaseOrders = []TicketNum{}
	err := buyer.Put(stub)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func GetBuyerByID(stub shim.ChaincodeStubInterface, userid *UserID) (*Buyer, error) {
	buyer := new(Buyer)
	key := userid.ToString() + "M"
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, buyer)
	if err != nil {
		return nil, err
	}
	return buyer, nil
}

func (u *Buyer) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(&u)
	if err != nil {
		return err
	}
	key := u.UserID.ToString() + "M"
	err = stub.PutState(key, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}

func (u *Buyer) AddPurchaseOrder(stub shim.ChaincodeStubInterface, ticketNum *TicketNum) error {
	u.PurchaseOrders = append(u.PurchaseOrders, *ticketNum)
	u.Put(stub)
	return nil
}

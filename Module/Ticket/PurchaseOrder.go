package Ticket

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

type PurchaseOrder struct {
	//order info
	TicketNumber       TicketNum               `json:"TicketNumber"`
	OrderDate          string                  `json:"OrderDate"`
	PurchaseProducts   []ProductPurchaseDetail `json:"PurchaseProducts"`
	OrderDepositAmount int64                   `json:"OrderDepositAmount"`
	TotalPrice         int64                   `json:"TotalPrice"`
	Status             int                     `json:"Status"`
	//buyer info
	Buyer                 UserID `json:"Buyer"`
	BuyerBank             UserID `json:"Buyerbank"`
	BuyerBankAccount      string `json:"BuyerBankAccount"`
	BuyerWarehouse        UserID `json:"BuyerWarehouse"`
	BuyerWarehouseAccount string `json:"BuyerWarehouseAccount"`
	//supplyer info
	Supplyer                 UserID    `json:"Supplyer"`
	SupplyerBank             UserID    `json:"SupplyerBank"`
	SupplyerBankAccount      string    `json:"SupplyerBankAccount"`
	SupplyerWarehouse        UserID    `json:"SupplyerWarehouse"`
	SupplyerWarehouseAccount string    `json:"SupplyerWarehouseAccount"`
	SupplyerConfirmDate      time.Time `json:"SupplyerConfirmDate"`
}

func (u *PurchaseOrder) InitPurchaseorder() PurchaseOrder {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return PurchaseOrder{}
}

func (u *PurchaseOrder) SetBuyerBank() error {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return nil
}

func (u *PurchaseOrder) SetBuyerWareHouse() error {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return nil
}

func (u *PurchaseOrder) AddPurchaseProduct() error {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return nil
}

func (u *PurchaseOrder) CalcTotalPrice() error {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	return nil
}

func (u *PurchaseOrder) SupplyConfirmWithAddBank() error {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	u.Status = CONFIRMED
	return nil
}

func (u *PurchaseOrder) GetPurchaseOrderByID(stub shim.ChaincodeStubInterface, purchaseOrderID TicketNum) (PurchaseOrder, error) {
	/*
	* TODO
	* 2017/2/25 luodanwg
	* */
	var po PurchaseOrder
	bytes, err := stub.GetState(purchaseOrderID.ToString())
	if err != nil {
		return PurchaseOrder{}, err
	}
	err = json.Unmarshal(bytes, &po)
	if err != nil {
		return PurchaseOrder{}, err
	}
	return po, nil
}

func (u *PurchaseOrder) SavePurchaseOrderByID(stub shim.ChaincodeStubInterface) error {
	/*
	* TODO Error catch
	* 2017/2/25 luodanwg
	* */
	jsonRespByte, _ := json.Marshal(&u)
	stub.PutState(u.TicketNumber.ToString(), jsonRespByte)
	return nil
}

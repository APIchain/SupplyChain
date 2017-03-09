package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Warehouse struct {
	UserID   UserID
	Capacity int32
	Address  string
	City     string
	Discount int8
	Account  []string
}

func NewWarehouseWithInit(stub shim.ChaincodeStubInterface, userid *UserID, cap int32, addr string, city string) (*Warehouse, error) {
	wh := new(Warehouse)
	wh.UserID = *userid
	wh.Capacity = cap
	wh.Address = addr
	wh.City = city
	wh.Discount = int8(100)
	wh.Put(stub)
	return wh, nil
}

func (u *Warehouse) SetDiscount(stub shim.ChaincodeStubInterface, discount int8) error {
	u.Discount = discount
	u.Put(stub)
	return nil
}

func GetWarehouseByID(stub shim.ChaincodeStubInterface, userid *UserID) (*Warehouse, error) {
	ubs := new(Warehouse)
	key := userid.ToString() + "W"
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, ubs)
	if err != nil {
		return nil, err
	}
	return ubs, nil
}

func (u *Warehouse) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, _ := json.Marshal(&u)
	key := u.UserID.ToString() + "W"
	stub.PutState(key, jsonRespByte)
	return nil
}

func (u *Warehouse) AddAccount(stub shim.ChaincodeStubInterface, warehouseID UserID, account string) error {
	u.Account = append(u.Account, account)
	u.Put(stub)
	return nil
}

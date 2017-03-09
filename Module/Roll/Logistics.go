package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Logistics struct {
	UserID        UserID `json:"UserID"`
	Discount      int8   `json:"Discount"`
	LogisticsBank Bank   `json:"LogisticsBank"`
}

func NewLogisticsWithInit(stub shim.ChaincodeStubInterface, userid *UserID) (*Logistics, error) {
	l := new(Logistics)
	l.UserID = *userid
	l.Discount = int8(100)
	return l, nil
}

func (u *Logistics) SetBank(stub shim.ChaincodeStubInterface, bank Bank) {
	u.LogisticsBank = bank
	u.Put(stub)
}

func (u *Logistics) SetDiscount(stub shim.ChaincodeStubInterface, discount int8) {
	u.Discount = discount
	u.Put(stub)
}

func GetLogisticsByID(stub shim.ChaincodeStubInterface, userid *UserID) (*Logistics, error) {
	ls := new(Logistics)
	key := userid.ToString() + "L"
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, ls)
	if err != nil {
		return nil, err
	}
	return ls, nil
}

func (u *Logistics) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, _ := json.Marshal(&u)
	key := u.UserID.ToString() + "L"
	stub.PutState(key, jsonRespByte)
	return nil
}

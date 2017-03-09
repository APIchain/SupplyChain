package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Supplyer struct {
	UserID            UserID      `json:"UserID"`
	Products          []ProductID `json:"Products"`
	DepositPayPercent int8        `json:"DepositPayPercent"`
}

func NewSupplyerWithInit(stub shim.ChaincodeStubInterface, userID *UserID) (*Supplyer, error) {
	supplyer := new(Supplyer)
	supplyer.UserID = *userID
	supplyer.Products = []ProductID{}
	supplyer.DepositPayPercent = 100
	supplyer.Put(stub)
	return supplyer, nil
}

func (u *Supplyer) AddProduct(stub shim.ChaincodeStubInterface, product *ProductID) error {
	u.Products = append(u.Products, *product)
	err := u.Put(stub)
	if err != nil {
		return err
	}
	return nil
}

func (u *Supplyer) SetDepositPayPercent(stub shim.ChaincodeStubInterface, depositPayPercent int8) error {
	u.DepositPayPercent = depositPayPercent
	err := u.Put(stub)
	if err != nil {
		return err
	}
	return nil
}

func GetSupplyerByID(stub shim.ChaincodeStubInterface, userid *UserID) (*Supplyer, error) {
	supplyer := new(Supplyer)
	key := userid.ToString() + "S"
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, supplyer)
	if err != nil {
		return nil, err
	}
	return supplyer, nil
}

func (u *Supplyer) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(&u)
	if err != nil {
		return err
	}
	key := u.UserID.ToString() + "S"
	err = stub.PutState(key, jsonRespByte)
	if err != nil {
		return err
	}
	return nil
}

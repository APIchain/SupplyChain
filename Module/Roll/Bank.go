package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Bank struct {
	UserID   UserID   `json:"UserID"`
	Accounts []string `json:"Accounts"`
}

func NewBankWithInit(stub shim.ChaincodeStubInterface, userid *UserID) (*Bank, error) {
	bank := new(Bank)
	bank.UserID = *userid
	bank.Accounts = []string{}
	err := bank.Put(stub)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func GetBankByID(stub shim.ChaincodeStubInterface, userid *UserID) (*Bank, error) {
	bank := new(Bank)
	key := userid.ToString() + "B"
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, bank)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func (u *Bank) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, _ := json.Marshal(&u)
	key := u.UserID.ToString() + "B"
	stub.PutState(key, jsonRespByte)
	return nil
}

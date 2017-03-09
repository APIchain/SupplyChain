package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type BusinessAccount struct {
	ID      UserID `json:"ID"`
	Account string `json:"Account"`
}

type UserBusiness struct {
	UserID        UserID            `json:"UserID"`
	ListBank      []BusinessAccount `json:"ListBank"`
	ListWarehouse []BusinessAccount `json:"ListWarehouse"`
}

func NewUserBusinessWithInit(stub shim.ChaincodeStubInterface, userid *UserID) (*UserBusiness, error) {
	ub := new(UserBusiness)
	ub.UserID = *userid
	ub.ListBank = []BusinessAccount{}
	ub.ListWarehouse = []BusinessAccount{}
	ub.Put(stub)
	return ub, nil
}

func (u *UserBusiness) AddBank(stub shim.ChaincodeStubInterface, bankid *UserID, account string) error {
	bs := BusinessAccount{
		ID:      *bankid,
		Account: account,
	}
	u.ListBank = append(u.ListBank, bs)
	u.Put(stub)
	return nil
}

func (u *UserBusiness) AddWarehouse(stub shim.ChaincodeStubInterface, warehouseID *UserID, account string) error {
	bs := BusinessAccount{
		ID:      *warehouseID,
		Account: account,
	}
	u.ListWarehouse = append(u.ListWarehouse, bs)
	u.Put(stub)
	return nil
}

func (u *UserBusiness) HasAccount(banid *UserID, account *string) bool {
	for _, v := range u.ListBank {
		if v.ID == *banid && v.Account == *account {
			return true
		}
	}
	return false
}

func GetUserBusinessByID(stub shim.ChaincodeStubInterface, userid *UserID) (*UserBusiness, error) {
	ubs := new(UserBusiness)
	key := userid.ToString() + "N"
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

func (u *UserBusiness) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, _ := json.Marshal(&u)
	key := u.UserID.ToString() + "N"
	stub.PutState(key, jsonRespByte)
	return nil
}

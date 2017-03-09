package Roll

import (
	. "Gyl/Module/System"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type UserDetail struct {
	Userid       UserID
	UserBase     *UserBase
	UserBusiness *UserBusiness
	Buyer        *Buyer
	Supplyer     *Supplyer
	Bank         *Bank
	Logistics    *Logistics
	Warehouse    *Warehouse
}

func NewUser(stub shim.ChaincodeStubInterface, username string) (*UserDetail, error) {
	user := new(UserDetail)
	userid, err := GenUserID(stub)
	if err != nil {
		return nil, err
	}
	userbase, err := NewUserBaseWithInit(stub, &user.Userid, username)
	if err != nil {
		return nil, err
	}
	return &UserDetail{
		Userid:       userid,
		UserBase:     userbase,
		UserBusiness: &UserBusiness{},
		Buyer:        &Buyer{},
		Supplyer:     &Supplyer{},
	}, nil
}

func (u *UserDetail) ToBeBuyer(stub shim.ChaincodeStubInterface) error {
	var err error
	u.UserBusiness, err = NewUserBusinessWithInit(stub, &u.Userid)
	if err != nil {
		return err
	}
	u.Buyer, err = NewBuyerWithInit(stub, &u.Userid)
	if err != nil {
		return err
	}
	u.UserBase.BeBuyer(stub)
	return nil
}

func (u *UserDetail) ToBeSupplyer(stub shim.ChaincodeStubInterface) error {
	var err error
	u.UserBusiness, err = NewUserBusinessWithInit(stub, &u.Userid)
	if err != nil {
		return err
	}
	u.Supplyer, err = NewSupplyerWithInit(stub, &u.Userid)
	if err != nil {
		return err
	}
	u.UserBase.BeSupplyer(stub)
	return nil
}

func (u *UserDetail) ToBeBanker(stub shim.ChaincodeStubInterface) error {
	var err error
	u.Bank, err = NewBankWithInit(stub, &u.Userid)
	if err != nil {
		return err
	}
	u.UserBase.BeBanker(stub)
	return nil
}

func (u *UserDetail) ToBeLogistics(stub shim.ChaincodeStubInterface) error {
	var err error
	u.Logistics, err = NewLogisticsWithInit(stub, &u.Userid)
	if err != nil {
		return err
	}
	u.UserBase.BeLogisticser(stub)
	return nil
}
func (u *UserDetail) ToBeWarehouse(stub shim.ChaincodeStubInterface, cap int32, addr string, city string) error {
	var err error
	u.Warehouse, err = NewWarehouseWithInit(stub, &u.Userid, cap, addr, city)
	if err != nil {
		return err
	}
	u.UserBase.BeWarehouser(stub)
	return nil
}

func GetUserDetailByID(stub shim.ChaincodeStubInterface, userid *UserID) (*UserDetail, error) {
	var err error
	u := new(UserDetail)
	u.Userid = *userid
	u.UserBase, err = GetUserBaseByID(stub, userid)
	if err != nil {
		return nil, err
	}
	u.UserBusiness, err = GetUserBusinessByID(stub, userid)
	if err != nil {
		return nil, err
	}
	u.Buyer, err = GetBuyerByID(stub, userid)
	if err != nil {
		return nil, err
	}
	u.Supplyer, err = GetSupplyerByID(stub, userid)
	if err != nil {
		return nil, err
	}
	u.Warehouse, err = GetWarehouseByID(stub, userid)
	if err != nil {
		return nil, err
	}
	u.Logistics, err = GetLogisticsByID(stub, userid)
	if err != nil {
		return nil, err
	}
	u.Bank, err = GetBankByID(stub, userid)
	if err != nil {
		return nil, err
	}
	return u, nil
}

package List

import (
	. "Gyl/Module/Roll"
	. "Gyl/Module/System"
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

const LIST_UserDetail = "ListUserDetailIDs"

var DefaultUserDetailLIst *ListUserDetailIDs

var DefaultUserAllDetailDetailList map[UserID]*UserDetail
var DefaultUserBuyerDetailDetailList map[UserID]*UserDetail
var DefaultUserSupplyerDetailDetailList map[UserID]*UserDetail
var DefaultUserWarehouseDetailDetailList map[UserID]*UserDetail
var DefaultUserBankDetailDetailList map[UserID]*UserDetail
var DefaultUserLogisticsDetailDetailList map[UserID]*UserDetail

type ListUserDetailIDs struct {
	UserDetailIDs []UserID `json:"UserDetailIDs"`
}

func InitDefaultUserList(stub shim.ChaincodeStubInterface) {
	if DefaultUserDetailLIst == nil {
		DefaultUserDetailLIst = new(ListUserDetailIDs)
		DefaultUserDetailLIst.get(stub)
		InitDefaultUserAllDetailDetailList(stub)
	}
}

func (l *ListUserDetailIDs) AddUserDetailIDs(stub shim.ChaincodeStubInterface, UserDetailid UserID) error {
	l.UserDetailIDs = append(l.UserDetailIDs, UserDetailid)
	l.put(stub)
	return nil
}

func (u *ListUserDetailIDs) get(stub shim.ChaincodeStubInterface) error {
	UserDetailIDs := new(ListUserDetailIDs)
	key := LIST_UserDetail
	bytes, err := stub.GetState(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, UserDetailIDs)
	if err != nil {
		return err
	}
	u = UserDetailIDs
	return nil
}

func (u *ListUserDetailIDs) put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, _ := json.Marshal(&u)
	key := LIST_UserDetail
	stub.PutState(key, jsonRespByte)
	return nil
}

func InitDefaultUserAllDetailDetailList(stub shim.ChaincodeStubInterface) error {
	for _, id := range DefaultUserDetailLIst.UserDetailIDs {
		UserDetailinfo, err := GetUserDetailByID(stub, &id)
		if err != nil {
			return err
		}
		DefaultUserAllDetailDetailList[id] = UserDetailinfo
		if UserDetailinfo.UserBase.IsBanker {
			DefaultUserBankDetailDetailList[id] = UserDetailinfo
		}
		if UserDetailinfo.UserBase.IsSupplyer {
			DefaultUserSupplyerDetailDetailList[id] = UserDetailinfo
		}
		if UserDetailinfo.UserBase.IsBuyer {
			DefaultUserBuyerDetailDetailList[id] = UserDetailinfo
		}
		if UserDetailinfo.UserBase.IsWareHouser {
			DefaultUserWarehouseDetailDetailList[id] = UserDetailinfo
		}
		if UserDetailinfo.UserBase.IsLogisticser {
			DefaultUserLogisticsDetailDetailList[id] = UserDetailinfo
		}
	}
	return nil

}

//func UpdOrSetDefaultUserDetailDetailList(stub shim.ChaincodeStubInterface, userid *UserID, UserDetail *UserDetail) {
//	DefaultUserAllDetailDetailList[*userid] = UserDetail
//}

func GetUserDetailDetailFromList(userid *UserID) (*UserDetail, error) {
	val, exist := DefaultUserAllDetailDetailList[*userid]
	if exist {
		return val, nil
	} else {
		return nil, errors.New("User ID not exist.id=" + userid.ToString())
	}
}

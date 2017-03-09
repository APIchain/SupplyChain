package Roll

import (
	. "Gyl/Module/System"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type UserBase struct {
	UserID        UserID `json:"UserID"`
	UserBaseName  string `json:"UserBaseName"`
	IsBuyer       bool   `json:"IsBuyer"`
	IsSupplyer    bool   `json:"IsSupplyer"`
	IsBanker      bool   `json:"IsBanker"`
	IsLogisticser bool   `json:IsLogisticser`
	IsWareHouser  bool   `json:IsWareHouser`
}

func NewUserBaseWithInit(stub shim.ChaincodeStubInterface, userid *UserID, username string) (*UserBase, error) {
	//创建实例
	user := new(UserBase)
	user.UserID = *userid
	user.UserBaseName = username
	user.IsBanker = false
	user.IsLogisticser = false
	user.IsWareHouser = false
	user.IsBuyer = false
	user.IsSupplyer = false

	//持久化
	err := user.Put(stub)
	if err != nil {
		Logger.Error("[Userbase] , Get failed with Put, key=." + userid.ToString())
		return nil, err
	}
	return user, nil
}

func GetUserBaseByID(stub shim.ChaincodeStubInterface, userid *UserID) (*UserBase, error) {
	userBase := new(UserBase)
	key := userid.ToString()
	bytes, err := stub.GetState(key)
	if err != nil {
		Logger.Error("[Userbase] , Get failed with GetState, key=." + key)
		return nil, err
	}
	err = json.Unmarshal(bytes, userBase)
	if err != nil {
		Logger.Error("[Userbase] , Get failed with Unmarshal, key=." + key)
		return nil, err
	}
	return userBase, nil
}

func (u *UserBase) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, err := json.Marshal(&u)
	if err != nil {
		Logger.Error("[Userbase] , Put failed with Unmarshal.")
	}
	stub.PutState(u.UserID.ToString(), jsonRespByte)
	return nil
}

func (u *UserBase) BeBuyer(stub shim.ChaincodeStubInterface) error {
	u.IsBuyer = true
	u.Put(stub)
	return nil
}

func (u *UserBase) BeBanker(stub shim.ChaincodeStubInterface) error {
	u.IsBanker = true
	u.Put(stub)
	return nil
}

func (u *UserBase) BeSupplyer(stub shim.ChaincodeStubInterface) error {
	u.IsSupplyer = true
	u.Put(stub)
	return nil
}

func (u *UserBase) BeLogisticser(stub shim.ChaincodeStubInterface) error {
	u.IsLogisticser = true
	u.Put(stub)
	return nil
}

func (u *UserBase) BeWarehouser(stub shim.ChaincodeStubInterface) error {
	u.IsWareHouser = true
	u.Put(stub)
	return nil
}

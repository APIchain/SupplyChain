package System

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

const USERID_KEY_COUNTER = "UserIDKeyCounter"

type UserID int64

func GenUserID(stub shim.ChaincodeStubInterface) (UserID, error) {
	key := USERID_KEY_COUNTER
	bytexs, err := stub.GetState(key)
	if err != nil {
		return UserID(uint64(0)), err
	}
	UserIDNow, err := strconv.ParseInt(string(bytexs), 10, 64)
	if err != nil {
		return UserID(uint64(0)), err
	}
	UserIDNow = UserIDNow + 1
	UserIDNowToWrite := strconv.FormatInt(UserIDNow, 10)
	if err != nil {
		return UserID(uint64(0)), err
	}
	buf := bytes.NewBufferString(UserIDNowToWrite)
	err = stub.PutState(key, buf.Bytes())
	if err != nil {
		return UserID(uint64(0)), err
	}
	temp := UserID(UserIDNow)
	return temp, nil
}

func (t *UserID) ToString() string {
	return fmt.Sprint(*t)
}

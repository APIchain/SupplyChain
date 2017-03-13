package main

import (
	"Gyl/Customer"
	"Gyl/List"
	shim "github.com/hyperledger/fabric/core/chaincode/shim"
)

//常量定义
const (
	CHAINCODE_LOG_FILE   = "gyl.log"
	CHAINCODE_LOG_PREFIX = ""
)

type Chaincode struct {
}

func (t *Chaincode) Initialize(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	logger.Infof("[WRITE LEDGER].KEY=[%s],VALUE=[%s]\n", key, string(putState_Byte))
	err = stub.PutState(key, putState_Byte)
	if err != nil {
		return nil, err
	}saa

	logger.Debugf("[WRITE LEDGER].KEY=[%s],VALUE=[%s]\n", key, string(putState_Byte))
	err = stub.PutState(key, putState_Byte)
	if err != nil {
		return nil, err
	}

	//init
	List.Init(stub)

	return nil, nil
}

//============================================================================================================
//     Function main函数
//============================================================================================================
func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		panic(err)
	}
}

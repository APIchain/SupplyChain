package System

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

const PRODUCT_KEY_COUNTER = "ProductKeyCounter"

type ProductID int64

func GenProductID(stub shim.ChaincodeStubInterface) (ProductID, error) {
	key := PRODUCT_KEY_COUNTER
	bytexs, err := stub.GetState(key)
	if err != nil {
		return ProductID(uint64(0)), err
	}
	ProductIDNow, err := strconv.ParseInt(string(bytexs), 10, 64)
	if err != nil {
		return ProductID(uint64(0)), err
	}
	ProductIDNow = ProductIDNow + 1
	ProductIDNowToWrite := strconv.FormatInt(ProductIDNow, 10)
	if err != nil {
		return ProductID(uint64(0)), err
	}
	buf := bytes.NewBufferString(ProductIDNowToWrite)
	err = stub.PutState(key, buf.Bytes())
	if err != nil {
		return ProductID(uint64(0)), err
	}
	temp := ProductID(ProductIDNow)
	return temp, nil
}

func (t *ProductID) ToString() string {
	return fmt.Sprint(*t)
}

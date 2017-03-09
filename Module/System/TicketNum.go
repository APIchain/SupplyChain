package System

import (
	"bytes"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

const TICKETNUM_KEY_COUNTER = "TicketNumKeyCounter"

type TicketNum int64

func GenTicketNum(stub shim.ChaincodeStubInterface) (TicketNum, error) {
	key := TICKETNUM_KEY_COUNTER
	bytexs, err := stub.GetState(key)
	if err != nil {
		return TicketNum(uint64(0)), err
	}
	TicketNumNow, err := strconv.ParseInt(string(bytexs), 10, 64)
	if err != nil {
		return TicketNum(uint64(0)), err
	}
	TicketNumNow = TicketNumNow + 1
	TicketNumNowToWrite := strconv.FormatInt(TicketNumNow, 10)
	if err != nil {
		return TicketNum(uint64(0)), err
	}
	buf := bytes.NewBufferString(TicketNumNowToWrite)
	err = stub.PutState(key, buf.Bytes())
	if err != nil {
		return TicketNum(uint64(0)), err
	}
	temp := TicketNum(TicketNumNow)
	return temp, nil
}

func (t *TicketNum) ToString() string {
	return fmt.Sprint(*t)
}

package System

import "github.com/hyperledger/fabric/core/chaincode/shim"

var Logger *shim.ChaincodeLogger

func init() {
	Logger = shim.NewLogger("gyl")
}

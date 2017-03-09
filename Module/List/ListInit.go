package List

import (
	"Gyl/Module/System"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

/**
 *
 * @author <a href="mailto:luodan.wg@gmail.com">junjie</a>
 * @version V1.0.0
 * @since 2017-02-27
 * Run init func to initialize the public var
 */
func Init(stub shim.ChaincodeStubInterface) {
	System.Logger.Debugf("[List] Init Started.\n")
	InitDefaultUserList(stub)
	InitListWaitingConfirm(stub)
	System.Logger.Debugf("[List] Init Completed.\n")
}

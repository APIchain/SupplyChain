package List

import (
	. "Gyl/Module/System"
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"time"
)

type ConfirmStatus int

const (
	CONFIRM_WAITING ConfirmStatus = iota
	CONFIRM_CONFIRMED
	CONFIRM_DECLINED
	//LOGISTICS_CONFIRM_WAITING
	//LOGISTICS_CONFIRM_CONFIRMED
	//LOGISTICS_CONFIRM_DECLINED
)

const LIST_WAITING_CONFIRM = "ListWaitingConfirm"

var StoListWaitingConfirm *ListWaitingConfirm

type ConfirmDetail struct {
	RequestedTime time.Time     `json:"RequestedTime"`
	ReplyedTime   time.Time     `json:"ReplyedTime"`
	WaitingID     UserID        `json:"WaitingID"`
	RequestedID   UserID        `json:"RequestedID"`
	TicketID      TicketNum     `json:"TicketID"`
	Status        ConfirmStatus `json:"Status"`
}

type ListWaitingConfirm struct {
	ConfirmDetails []ConfirmDetail `json:"ConfirmDetail"`
}

func InitListWaitingConfirm(stub shim.ChaincodeStubInterface) {
	if StoListWaitingConfirm == nil {
		StoListWaitingConfirm = new(ListWaitingConfirm)
		StoListWaitingConfirm.Get(stub)
	}
}

func (l *ListWaitingConfirm) AddToList(stub shim.ChaincodeStubInterface, waitingid UserID, fromid UserID, tickid TicketNum) error {

	for _, v := range l.ConfirmDetails {
		if v.TicketID == tickid && v.WaitingID == waitingid && v.Status == CONFIRM_WAITING {
			return errors.New("This Ticket has already Exist in Confirm Waiting list,Please Waiting for process")
		}
	}
	cd := new(ConfirmDetail)
	cd.RequestedTime = time.Now()
	cd.WaitingID = waitingid
	cd.RequestedID = fromid
	cd.TicketID = tickid
	l.ConfirmDetails = append(l.ConfirmDetails, *cd)
	l.Put(stub)
	return nil
}

func (u *ListWaitingConfirm) Get(stub shim.ChaincodeStubInterface) error {
	var lcf ListWaitingConfirm
	key := LIST_WAITING_CONFIRM
	bytes, err := stub.GetState(key)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, &lcf)
	if err != nil {
		return err
	}
	u = &lcf
	return nil
}

func (u *ListWaitingConfirm) Put(stub shim.ChaincodeStubInterface) error {
	jsonRespByte, _ := json.Marshal(&u)
	key := LIST_WAITING_CONFIRM
	stub.PutState(key, jsonRespByte)
	return nil
}

func (u *ListWaitingConfirm) GetWithWaitingID(stub shim.ChaincodeStubInterface, waitingid *UserID) (*[]ConfirmDetail, error) {
	confirmDetails := []ConfirmDetail{}
	for _, v := range u.ConfirmDetails {
		if v.WaitingID == *waitingid {
			confirmDetails = append(confirmDetails, v)
		}
	}
	return &confirmDetails, nil
}

func (u *ListWaitingConfirm) GetWithRequestedID(stub shim.ChaincodeStubInterface, requesintedID UserID) (*[]ConfirmDetail, error) {
	confirmDetails := []ConfirmDetail{}
	for _, v := range u.ConfirmDetails {
		if v.RequestedID == requesintedID {
			confirmDetails = append(confirmDetails, v)
		}
	}
	return &confirmDetails, nil
}

func (u *ListWaitingConfirm) Confirm(stub shim.ChaincodeStubInterface, ticketid TicketNum, approver UserID, newstatus ConfirmStatus) error {
	for _, v := range u.ConfirmDetails {
		if v.TicketID == ticketid && v.WaitingID == approver && v.Status == CONFIRM_WAITING {
			v.Status = CONFIRM_CONFIRMED
			u.Put(stub)
			return nil
		} else {
			return errors.New("Please check the TickitID or the Tickit's status.")
		}
	}
	return nil
}

func (u *ListWaitingConfirm) Decline(stub shim.ChaincodeStubInterface, ticketid TicketNum, approver UserID, newstatus ConfirmStatus) error {
	for _, v := range u.ConfirmDetails {
		if v.TicketID == ticketid && v.WaitingID == approver && v.Status == CONFIRM_WAITING {
			v.Status = CONFIRM_DECLINED
			u.Put(stub)
			return nil
		} else {
			return errors.New("Please check the TickitID or the Tickit's status.")
		}
	}
	return nil
}

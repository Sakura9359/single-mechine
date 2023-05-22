package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
	"time"
)

func AddRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	fmt.Println("- start add request record")
	requester := args[0]
	role := args[1]
	pubKey := args[2]
	dataID := args[3]
	dataOwner := args[4]
	level, _ := strconv.Atoi(args[5])
	requestID := args[6]
	sa := SA{requester, role, pubKey}
	da := DA{dataID, dataOwner, dataID}
	txTime, _ := stub.GetTxTimestamp()
	timestamp := fmt.Sprint(txTime)
	reqRecord := &RequestRecord{requestID, sa, da, level, timestamp}
	reqRecordJSONasBytes, err := json.Marshal(reqRecord)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(requestID, reqRecordJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add request record")
	return shim.Success([]byte(requestID))
}

func AddResponse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	fmt.Println("- start add response record")
	owner := args[0]
	endtime := args[1]
	requestID := args[2]
	responseID := args[3]
	policyID := args[4]

	timeObj, err := time.Parse("2006-01-02 15:04:05", endtime)
	endtime = strconv.Itoa(int(timeObj.UnixNano()))
	txTime, _ := stub.GetTxTimestamp()
	timestamp := fmt.Sprint(txTime)
	responseRecord := &ResponseRecord{responseID, policyID, owner, requestID, endtime, timestamp}
	reqRecordJSONasBytes, err := json.Marshal(responseRecord)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(responseID, reqRecordJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add response record")
	return shim.Success([]byte(responseID))
}

func AddAction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	fmt.Println("- start add action record")
	userID := args[0]
	dataID := args[1]
	action := args[2]
	policyID := args[3]
	recordID := args[4]
	txTime, _ := stub.GetTxTimestamp()
	timestamp := fmt.Sprint(txTime)
	acRecord := &ActionRecord{recordID, policyID, userID, dataID, action, timestamp}
	acRecordJSONasBytes, err := json.Marshal(acRecord)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(recordID, acRecordJSONasBytes) // key = hash(userID + dataID + timestamp)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add action record")
	return shim.Success([]byte(recordID))
}

func GetRequest(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start get request record")
	requestID := args[0]
	reqRecordAsbytes, err := stub.GetState(requestID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	return shim.Success(reqRecordAsbytes)
}

func GetResponse(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start get response record")
	responseID := args[0]
	resRecordAsbytes, err := stub.GetState(responseID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	return shim.Success(resRecordAsbytes)
}

func GetAcRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start get action record")
	RecordID := args[0]
	acRecordAsbytes, err := stub.GetState(RecordID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	return shim.Success(acRecordAsbytes)
}

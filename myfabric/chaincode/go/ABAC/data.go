package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
	"time"
)

func InitData(stub shim.ChaincodeStubInterface) pb.Response {
	dataID := "豫E-MJ893"
	cid := "486415364153asd486aw4d68wa41d36aw4d"

	txTime, _ := stub.GetTxTimestamp()
	timeNow := fmt.Sprint(txTime)
	data := &Data{dataID, cid, timeNow, timeNow}
	dataJSONasBytes, err := json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dataID, dataJSONasBytes) // key := dataID
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func AddData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	fmt.Println("- start add data")
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}

	dataID := args[0]
	cid := args[1]

	dataAsbytes, err := stub.GetState(dataID)
	if err != nil {
		return shim.Error("Failed to get car record:" + err.Error())
	} else if dataAsbytes != nil {
		fmt.Println("This car record already exists:" + dataID)
		return shim.Error("This car record already exists:" + dataID)
	}
	txTime, _ := stub.GetTxTimestamp()
	timeNow := fmt.Sprint(txTime)
	data := &Data{dataID, cid, timeNow, timeNow}
	dataJSONasBytes, err := json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dataID, dataJSONasBytes) // key := dataID
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add data")
	return shim.Success(nil)
}

func UpdateData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}

	fmt.Println("- start update data")
	userID := args[0]
	dataID := args[1]
	cid := args[2]
	policyID := args[3]
	actionID := args[4]
	policyAsbytes, err := stub.GetState(policyID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	ap := &AP{}
	json.Unmarshal(policyAsbytes, ap)
	// 验证数字签名
	policy := ap.Policy
	if !Verify(ap) {
		return shim.Error("policy has been changed")
	}
	_, _, c, _, err := ActionParse(policy.ActionA.Level)
	if err != nil {
		return shim.Error(err.Error())
	} else if c != 1 {
		return shim.Error("permission no")
	}

	dataAsbytes, err := stub.GetState(dataID)
	if err != nil {
		return shim.Error("Failed to get car record:" + err.Error())
	} else if dataAsbytes == nil {
		fmt.Println("This car record does not exists:" + dataID)
		return AddData(stub, []string{dataID, cid})
	}
	txTime, _ := stub.GetTxTimestamp()
	timeNow := fmt.Sprint(txTime)
	data := &Data{}
	err = json.Unmarshal(dataAsbytes, data)
	if err != nil {
		return shim.Error(err.Error())
	}
	data.UpdateTime = timeNow
	data.Cid = cid
	dataJSONasBytes, err := json.Marshal(data)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(dataID, dataJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	// add action record
	argss := []string{userID, dataID, "UpdateData", policyID, actionID}
	AddAction(stub, argss)
	fmt.Println("- end update data")
	return shim.Success(nil)
}

func GetData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	fmt.Println("- start get data")
	userID := args[0]
	dataID := args[1]
	policyID := args[2]
	actionID := args[3]
	policyAsbytes, err := stub.GetState(policyID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	ap := &AP{}
	json.Unmarshal(policyAsbytes, ap)
	// 验证数字签名
	policy := ap.Policy
	if !Verify(ap) {
		return shim.Error("policy has been changed")
	}
	_, _, _, d, err := ActionParse(policy.ActionA.Level)
	if err != nil {
		return shim.Error(err.Error())
	} else if d != 1 {
		return shim.Error("permission no")
	}

	dataAsbytes, err := stub.GetState(dataID)
	if err != nil {
		return shim.Error("Failed to get car record:" + err.Error())
	} else if dataAsbytes == nil {
		fmt.Println("This car record does not exists:" + dataID)
		return shim.Error("This car record does not exists:" + dataID)
	}
	// add action record
	argss := []string{userID, dataID, "GetData", policyID, actionID}
	AddAction(stub, argss)

	fmt.Println("- end get data")
	return shim.Success(dataAsbytes)
}

func DeleteData(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	userID := args[0]
	dataID := args[1]
	policyID := args[2]
	actionID := args[3]
	policyAsbytes, err := stub.GetState(policyID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	ap := &AP{}
	json.Unmarshal(policyAsbytes, ap)
	// 验证数字签名
	policy := ap.Policy
	if !Verify(ap) {
		return shim.Error("policy has been changed")
	}
	_, b, _, _, err := ActionParse(policy.ActionA.Level)
	if err != nil {
		return shim.Error(err.Error())
	} else if b != 1 {
		return shim.Error("permission no")
	}

	err = stub.DelState(dataID)
	if err != nil {
		return shim.Error("Failed to delete data:" + err.Error())
	}
	// add action record
	argss := []string{userID, dataID, "DeleteData", policyID, actionID}
	AddAction(stub, argss)
	return shim.Success(nil)
}

func GetHistoryRecord(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	userID := args[0]
	dataID := args[1]
	policyID := args[2]
	actionID := args[3]
	policyAsbytes, err := stub.GetState(policyID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	ap := &AP{}
	json.Unmarshal(policyAsbytes, ap)
	// 验证数字签名
	policy := ap.Policy
	if !Verify(ap) {
		return shim.Error("policy has been changed")
	}
	a, _, _, _, err := ActionParse(policy.ActionA.Level)
	if err != nil {
		return shim.Error(err.Error())
	} else if a != 1 {
		return shim.Error("permission no")
	}

	fmt.Printf("- start getHistoryRecord: %s\n", dataID)
	resultsIterator, err := stub.GetHistoryForKey(dataID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the marble
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")

		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	// add action record
	argss := []string{userID, dataID, "GetHistoryRecord", policyID, actionID}
	AddAction(stub, argss)
	fmt.Printf("- getHistoryRecord returning:\n%s\n", buffer.String())
	return shim.Success(buffer.Bytes())
}

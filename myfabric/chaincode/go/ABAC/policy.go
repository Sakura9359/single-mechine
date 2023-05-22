package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

func AddPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	fmt.Println("- start add policy")

	policyID := args[0]
	requestID := args[1]
	responseID := args[2]
	policyStr := args[3]
	signature := args[4]

	policy := Policy{}
	json.Unmarshal([]byte(policyStr), &policy)

	ap := &AP{policyID, requestID, responseID, policy, signature}
	policyJSONasBytes, err := json.Marshal(ap)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(policyID, policyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add policy")
	return shim.Success(nil)
}

func MatchPolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start match policy")
	policyID := args[0]
	policyAsbytes, err := stub.GetState(policyID)
	if err != nil {
		return shim.Error("Failed to get policy info:" + err.Error())
	}
	return shim.Success(policyAsbytes)
}

func UpdatePolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	fmt.Println("- start add policy")

	policyID := args[0]
	requestID := args[1]
	responseID := args[2]
	policyStr := args[3]
	signature := args[4]

	policy := Policy{}
	json.Unmarshal([]byte(policyStr), &policy)

	ap := &AP{policyID, requestID, responseID, policy, signature}
	policyJSONasBytes, err := json.Marshal(ap)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(policyID, policyJSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("- end add policy")
	return shim.Success(nil)
}

func DeletePolicy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	fmt.Println("- start delete policy")
	policyID := args[0]
	err := stub.DelState(policyID)
	if err != nil {
		return shim.Error("Failed to delete policy:" + err.Error())
	}
	return shim.Success(nil)
}

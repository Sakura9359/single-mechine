package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

type SmartContract struct {
}

func (sc *SmartContract) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// 初始化用户
	args := stub.GetStringArgs()
	user1SK, user1PK := decode(args[0], args[1])
	user2SK, user2PK := decode(args[2], args[3])
	user1 := &User{"user1", "豫E-MJ893", "org1", "dep1", "nil", user1PK, user1SK}
	user2 := &User{"user2", "豫E-MJ893", "org1", "dep1", "nil", user2PK, user2SK}
	user1JSONasBytes, _ := json.Marshal(user1)
	user2JSONasBytes, _ := json.Marshal(user2)
	err := stub.PutState(user1.UserID, user1JSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(user2.UserID, user2JSONasBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	// 初始化数据
	InitData(stub)
	return shim.Success(nil)
}

func (sc *SmartContract) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoke is running " + function)
	if function == "AddData" {
		return AddData(stub, args)
	} else if function == "UpdateData" {
		return UpdateData(stub, args)
	} else if function == "GetData" {
		return GetData(stub, args)
	} else if function == "DeleteData" {
		return DeleteData(stub, args)
	} else if function == "GetHistoryRecord" {
		return GetHistoryRecord(stub, args)
	} else if function == "AddPolicy" {
		return AddPolicy(stub, args)
	} else if function == "MatchPolicy" {
		return MatchPolicy(stub, args)
	} else if function == "UpdatePolicy" {
		return UpdatePolicy(stub, args)
	} else if function == "DeletePolicy" {
		return DeletePolicy(stub, args)
	} else if function == "AddRequest" {
		return AddRequest(stub, args)
	} else if function == "AddResponse" {
		return AddResponse(stub, args)
	} else if function == "AddAction" {
		return AddAction(stub, args)
	} else if function == "GetRequest" {
		return GetRequest(stub, args)
	} else if function == "GetResponse" {
		return GetResponse(stub, args)
	} else if function == "GetAcRecord" {
		return GetAcRecord(stub, args)
	}

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

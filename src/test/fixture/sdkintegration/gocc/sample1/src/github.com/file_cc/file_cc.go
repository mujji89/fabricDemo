/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	//"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("fileBlockchain-logger")

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

// Init initializes the chaincode state
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc Init ###########")
	_, args := stub.GetFunctionAndParameters()

	var fileName, contractName, owner, accessList string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fileName = args[0]
	contractName = args[1]
	owner = args[2]
	accessList = args[3]
	jsonResp := "{\"fileHash\":\"" + fileName + "\",\"owner\":\"" + owner + "\",\"accesslist\":\"" + accessList + "\"}"


	err = stub.PutState(contractName, []byte(jsonResp))
	if err != nil {
		return shim.Error(err.Error())
	}


	/*	_, args := stub.GetFunctionAndParameters()
		var A, B string    // Entities
		var Aval, Bval int // Asset holdings
		var err error

		if len(args) != 4 {
			return shim.Error("Incorrect number of arguments. Expecting 4")
		}

		// Initialize the chaincode
		A = args[0]
		Aval, err = strconv.Atoi(args[1])
		if err != nil {
			return shim.Error("Expecting integer value for asset holding")
		}
		B = args[2]
		Bval, err = strconv.Atoi(args[3])
		if err != nil {
			return shim.Error("Expecting integer value for asset holding")
		}
		logger.Infof("Aval = %d, Bval = %d\n", Aval, Bval)

		// Write the state to the ledger
		err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
		if err != nil {
			return shim.Error(err.Error())
		}

		err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
		if err != nil {
			return shim.Error(err.Error())
		}

		if transientMap, err := stub.GetTransient(); err == nil {
			if transientData, ok := transientMap["result"]; ok {
				return shim.Success(transientData)
			}
		}*/
	return shim.Success(nil)

}

// Invoke makes payment of X units from A to B
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	logger.Info("########### example_cc Invoke ###########")

	function, args := stub.GetFunctionAndParameters()

	if function == "delete" {
		// Deletes an entity from its state
		return t.delete(stub, args)
	}

	if function == "query" {
		// queries an entity state
		return t.query(stub, args)
	}

	if function == "add" {
		// Deletes an entity from its state
		return t.add(stub, args)
	}

	logger.Errorf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0])
	return shim.Error(fmt.Sprintf("Unknown action, check the first argument, must be one of 'delete', 'query', or 'move'. But got: %v", args[0]))
}

func (t *SimpleChaincode) add(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// must be an invoke
	logger.Info("########### example_cc Invoke ###########")


	var fileName, contractName, owner, accessList string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}

	fileName = args[0]
	contractName = args[1]
	owner = args[2]
	accessList = args[3]
	jsonResp := "{\"fileHash\":\"" + fileName + "\",\"owner\":\"" + owner + "\",\"accesslist\":\"" + accessList + "\"}"


	err = stub.PutState(contractName, []byte(jsonResp))
	if err != nil {
		return shim.Error(err.Error())
	}

/*	// Get the state from the ledger
	// TODO: will be nice to have a GetAllState call to ledger
	Avalbytes, err := stub.GetState(A)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Avalbytes == nil {
		return shim.Error("Entity not found")
	}
	Aval, _ = strconv.Atoi(string(Avalbytes))

	Bvalbytes, err := stub.GetState(B)
	if err != nil {
		return shim.Error("Failed to get state")
	}
	if Bvalbytes == nil {
		return shim.Error("Entity not found")
	}
	Bval, _ = strconv.Atoi(string(Bvalbytes))

	// Perform the execution
	X, err = strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expecting a integer value")
	}
	Aval = Aval - X
	Bval = Bval + X
	logger.Infof("Aval = %d, Bval = %d\n", Aval, Bval)

	// Write the state back to the ledger
	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
	if err != nil {
		return shim.Error(err.Error())
	}

	if transientMap, err := stub.GetTransient(); err == nil {
		if transientData, ok := transientMap["event"]; ok {
			stub.SetEvent("event", transientData)
		}
		if transientData, ok := transientMap["result"]; ok {
			return shim.Success(transientData)
		}
	}*/
	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	A := args[0]

	// Delete the key from the state in ledger
	err := stub.DelState(A)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	var contractName string    // Entities
	//var Aval, Bval int // Asset holdings
	//var X int          // Transaction value
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1 contract name")
	}


	//fileName = args[0]
	contractName = args[0]
	//requesterKey = args[3];

	// Get the state from the ledger
	Avalbytes, err := stub.GetState(contractName)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to get state for " + contractName + "\"}"
		return shim.Error(jsonResp)
	}

	if Avalbytes == nil {
		jsonResp := "{\"Error\":\"Nil amount for " + contractName + "\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"contractName\":\"" + contractName + "\",\"Values\":\"" + string(Avalbytes) + "\"}"
	logger.Infof("Query Response:%s\n", jsonResp)
	return shim.Success(Avalbytes)
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		logger.Errorf("Error starting Simple chaincode: %s", err)
	}
}

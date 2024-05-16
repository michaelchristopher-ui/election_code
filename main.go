package main

import (
	"election_code/smart-contract/core"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// The main function of the Election Chaincode
func main() {
	cc, err := contractapi.NewChaincode(&core.ElectionChainCode{})

	if err != nil {
		panic(err.Error())
	}

	if err := cc.Start(); err != nil {
		panic(err.Error())
	}
}

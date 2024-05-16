package helpers

import (
	"election_code/smart-contract/errors"
	"election_code/smart-contract/structs"
	"election_code/smart-contract/structs/fun"
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GetElectionDataInternal is a function that gets the election data from the world state and returns it within an election struct.
func GetElectionDataInternal(ctx contractapi.TransactionContextInterface, req fun.GetElectionDataReq) (structs.Election, error) {
	election, err := ctx.GetStub().GetState(req.Key)
	if err != nil {
		return structs.Election{}, err
	}
	if election == nil {
		return structs.Election{}, errors.ErrElectionNotExist(req.Key)
	}

	var electionStruct structs.Election
	err = json.Unmarshal(election, &electionStruct)
	if err != nil {
		return structs.Election{}, err
	}
	return electionStruct, nil
}

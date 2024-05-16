package helpers

import (
	"election_code/smart-contract/errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// GetState is a wrapper function for the GetState function of the interface that gets the value byte array from the world state
func GetState(ctx contractapi.TransactionContextInterface, key string) ([]byte, error) {
	data, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.ErrDataDoesNotExist
	}
	return data, nil
}

// CheckIfExists is a function that checks if a key and a value pair exists within the world state.
// It returns an error when the data returned is nil, which indicates that the key value pair does not exist.
// Code that uses this must treat the error return as if the data does not exist.
func CheckIfExists(ctx contractapi.TransactionContextInterface, key string) error {
	data, err := ctx.GetStub().GetState(key)
	if err != nil {
		return err
	}
	if data != nil {
		return errors.ErrDataExists
	}
	return nil
}

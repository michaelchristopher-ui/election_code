package helpers

import (
	"election_code/smart-contract/errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// DeleteState is a wrapper function for the DeleteState function of the interface, which deletes a key value pair from the world state
func DeleteState(ctx contractapi.TransactionContextInterface, key string) error {
	err := ctx.GetStub().DelState(key)
	if err != nil {
		return errors.ErrDataDoesNotExist
	}
	return nil
}

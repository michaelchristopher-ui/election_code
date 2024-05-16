package helpers

import (
	"election_code/smart-contract/errors"
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// PutState marshals the value and puts the resulting byte array into the world state with the supplied key string as the key.
// The interface value must be able to be marshalled using json.Marshal. This would typically be a json tagged struct.
func PutState(ctx contractapi.TransactionContextInterface, key string, value interface{}) error {
	valueByte, err := json.Marshal(value)
	if err != nil {
		return err
	}
	err = ctx.GetStub().PutState(key, valueByte)
	if err != nil {
		return errors.ErrPutState("PutState", key)
	}
	return nil
}

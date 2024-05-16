package helpers

import (
	"election_code/smart-contract/constants"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// IsRole checks whether the identity has the attribute and whether it is equal to the supplied value.
func IsRole(ctx contractapi.TransactionContextInterface, attrName constants.RoleString, attrValue string) bool {
	return ctx.GetClientIdentity().AssertAttributeValue(string(attrName), attrValue) == nil
}

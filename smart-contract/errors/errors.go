package errors

import (
	"election_code/smart-contract/constants"
	"errors"
	"fmt"
)

var (
	//Format texts for errors. Do not use this directly, call the wrapper function for it.
	//When creating a new error with a format, create the text then create the wrapper function.
	MarshalFailureText      = "failure when marshaling on function %s"
	PutStateText            = "failure when putting state on function %s, key %s"
	GetStateText            = "failure when getting state on function %s, key %s"
	ACLText                 = "failure when executing function %s, user does not have role %s"
	ElectionNotExistText    = "election with key %s does not exist"
	BallotNotExistText      = "ballot with key %s does not exist"
	VoterNotExistText       = "voter with key %s does not exist"
	VotableItemNotExistText = "votable item with key %s does not exist"

	//Errors without formatting.
	ErrAlreadyVoted     = errors.New("user has voted for this ballot")
	ErrDataDoesNotExist = errors.New("data does not exist")
	ErrDataExists       = errors.New("data already exists")
	ErrNotElectionTime  = errors.New("it's not election time")
	ErrBallotNotCreated = errors.New("ballot has not been created for this user")
)

// Wrapper for the errors
func ErrMarshalFailure(functionName string) error {
	return fmt.Errorf(MarshalFailureText, functionName)
}

func ErrPutState(functionName string, key string) error {
	return fmt.Errorf(PutStateText, functionName, key)
}

func ErrGetState(functionName string, key string) error {
	return fmt.Errorf(GetStateText, functionName, key)
}

func ErrACL(functionName string, roleString constants.RoleString) error {
	return fmt.Errorf(ACLText, functionName, roleString)
}

func ErrElectionNotExist(key string) error {
	return fmt.Errorf(ElectionNotExistText, key)
}

func ErrBallotNotExist(key string) error {
	return fmt.Errorf(BallotNotExistText, key)
}

func ErrVoterNotExist(key string) error {
	return fmt.Errorf(VoterNotExistText, key)
}

func ErrVotableItemNotExist(key string) error {
	return fmt.Errorf(VotableItemNotExistText, key)
}

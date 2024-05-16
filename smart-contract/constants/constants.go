package constants

//RoleString is a custom string type for id attribute keys that mark an invoker's role.
type RoleString string

const (
	//Custom Date format for any date within the whole chaincode
	DateFormat = "2006-01-02"

	//Constant values for id attribute keys
	AdminRole RoleString = "election.admin"
	VoterRole RoleString = "election.voter"

	//String versions of booleans
	TrueString  = "true"
	FalseString = "false"
)

package structs

//Voter is the struct representation of the Voter data
type Voter struct {
	VoterId     string
	RegistrarId string
	FirstName   string
	LastName    string
	BallotVoted map[string]bool
}

//NewVoterReq is the parameter struct for NewVoter
type NewVoterReq struct {
	VoterId     string
	RegistrarId string
	FirstName   string
	LastName    string
}

//NewVoteris the constructor of Voter
func NewVoter(req NewVoterReq) Voter {
	return Voter{
		VoterId:     req.VoterId,
		RegistrarId: req.RegistrarId,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		BallotVoted: map[string]bool{},
	}
}

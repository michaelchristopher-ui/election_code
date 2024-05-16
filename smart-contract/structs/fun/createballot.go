package fun

//CreateBallotReq is the request json struct of the CreateBallot function
type CreateBallotReq struct {
	ElectionId string `json:"election_id"`
	BallotCast string `json:"ballot_cast"`
	BallotId   string `json:"ballot_id"`
}

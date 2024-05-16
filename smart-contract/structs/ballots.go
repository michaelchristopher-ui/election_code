package structs

//Ballot is the struct representation of the Ballot data
type Ballot struct {
	VotableItems map[string]VotableItem `json:"votable_items"` //The key is the VotableItemId
	VoterId      string                 `json:"voter_id"`
	BallotCast   string                 `json:"ballot_cast"`
	BallotId     string                 `json:"ballot_id"`
}

//NewBallotReq is the parameter struct for the Ballot Constructor
type NewBallotReq struct {
	VotableItems map[string]VotableItem
	VoterId      string
	BallotCast   string
	BallotId     string
}

//NewBallot is the constructor for Ballot
func NewBallot(req NewBallotReq) Ballot {
	return Ballot{
		VotableItems: req.VotableItems,
		VoterId:      req.VoterId,
		BallotCast:   req.BallotCast,
		BallotId:     req.BallotId,
	}
}

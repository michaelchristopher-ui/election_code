package fun

//CreateVotableItemsReq is the request json struct of the CreateVotableItems function
type CreateVotableItemsReq struct {
	VotableId     string `json:"votable_id"`
	Description   string `json:"description"`
	Count         int    `json:"count"`
	ElectionIndex string `json:"election_index"`
	BallotIndex   string `json:"ballot_index"`
}

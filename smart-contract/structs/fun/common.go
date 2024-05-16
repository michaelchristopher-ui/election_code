package fun

//VotableItemRes is the return json struct representation of the VotableItems struct
type VotableItemRes struct {
	VotableId   string `json:"votable_id"`
	Description string `json:"description"`
	Count       int    `json:"count,omitempty"`
}

//BallotRes is the return json struct representation of the Ballot Struct
type BallotRes struct {
	VotableItems []VotableItemRes
	VoterId      string
	BallotCast   string
	BallotId     string
}

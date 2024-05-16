package structs

//Election is the struct representation of the Election data
type Election struct {
	ElectionId string            `json:"election_id"`
	Name       string            `json:"name"`
	Country    string            `json:"country"`
	Year       string            `json:"year"`
	StartDate  string            `json:"start_date"`
	EndDate    string            `json:"end_date"`
	Ballots    map[string]Ballot `json:"ballots"` //The key is the BallotId
}

//NewElectionReq is the parameter struct for the Election constructor
type NewElectionReq struct {
	ElectionId string
	Name       string
	Country    string
	Year       string
	StartDate  string
	EndDate    string
}

//NewElection is the constructor for Election
func NewElection(req NewElectionReq) Election {
	return Election{
		ElectionId: req.ElectionId,
		Name:       req.Name,
		Country:    req.Year,
		StartDate:  req.StartDate,
		EndDate:    req.EndDate,
	}
}

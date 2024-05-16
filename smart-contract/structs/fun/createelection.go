package fun

//CreateElectionReq is the request json struct of the CreateElection function
type CreateElectionReq struct {
	ElectionId string `json:"election_id"`
	Name       string `json:"name"`
	Country    string `json:"country"`
	Year       string `json:"year"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

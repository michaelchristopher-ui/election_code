package fun

//GetElectionDataReq is a request json struct for the GetElectionData function
type GetElectionDataReq struct {
	Key string `json:"key"`
}

//GetElectionDataRes is a return json struct for the GetElectionData function
type GetElectionDataRes struct {
	ElectionId string               `json:"election_id"`
	Name       string               `json:"name"`
	Country    string               `json:"country"`
	Year       string               `json:"year"`
	StartDate  string               `json:"start_date"`
	EndDate    string               `json:"end_date"`
	Ballots    map[string]BallotRes `json:"votables"`
}

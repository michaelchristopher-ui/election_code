package fun

//CreateVoterReq is a request json struct for the CreateVoter function
type CreateVoterReq struct {
	VoterId     string `json:"voter_id"`
	RegistrarId string `json:"registrar_id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

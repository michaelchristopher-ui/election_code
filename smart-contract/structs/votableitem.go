package structs

//VotableItem is the struct representation of the Votable Item data
type VotableItem struct {
	VotableId   string `json:"votable_id"`
	Description string `json:"description"`
	Count       int    `json:"count"`
}

//NewVotableItem is the constructor for VotableItem
func NewVotableItem(votableId string, description string) VotableItem {
	return VotableItem{
		VotableId:   votableId,
		Description: description,
		Count:       0,
	}
}

package dto

type CreateMemberDTO struct {
	Name      string `json:"name" validate:"empty=false"`
	Telephone string `json:"telephone" validate:"empty=false"`
}

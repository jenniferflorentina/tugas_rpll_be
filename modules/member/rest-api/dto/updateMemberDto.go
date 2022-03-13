package dto

type UpdateMemberDTO struct {
	Name      string `json:"name"`
	Point     int64  `json:"point"`
	Telephone string `json:"telephone"`
}

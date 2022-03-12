package dto

type CreateUpdateProductCategoryDTO struct {
	Name string `json:"name" validate:"empty=false"`
}

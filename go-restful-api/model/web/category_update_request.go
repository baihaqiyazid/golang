package web

type CategoryUpdateRequest struct{
	Id int `validate:"required"`
	Name string `validate:"required,lte=50,gte=1" json:"name"`
}
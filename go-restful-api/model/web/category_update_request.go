package web

type CategoryUpdateRequest struct{
	Id int `validate:"required"`
	Name string `validate:"required, max=50, min=1"`
}
package web

type CategoryCreateRequest struct{
	Name string `validate:"required, max=50, min=1"`
}
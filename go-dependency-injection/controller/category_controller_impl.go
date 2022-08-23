package controller

import (
	"go-restful-api/helper"
	"go-restful-api/model/web"
	"go-restful-api/services"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct{
	CategoryService services.CategoryService
}

func NewCategoryController(categoryService services.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)
	
	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success",
		Data: categoryResponse,
	}

	helper.WriteToWebResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)  {
	
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)
	
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.Panic(err)

	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success",
		Data: categoryResponse,
	}

	helper.WriteToWebResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.Panic(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success",
	}

	helper.WriteToWebResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) GetById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	
	categoryId := params.ByName("categoryId")
	id, err := strconv.Atoi(categoryId)
	helper.Panic(err)

	categoryResponse := controller.CategoryService.GetById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success",
		Data: categoryResponse,
	}

	helper.WriteToWebResponse(writer, webResponse)
}

func (controller *CategoryControllerImpl) GetAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponses := controller.CategoryService.GetAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "Success",
		Data: categoryResponses,
	}

	helper.WriteToWebResponse(writer, webResponse)
}
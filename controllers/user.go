package controllers

import (
	"GoDEMO/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserDTO struct {
	Name string `json:"username"`
	Age  int    `json:"age"`
	Id   string `json:"id"`
}

type IUserController interface {
	Add(context *gin.Context)
	GetById(context *gin.Context)
	GetAll(context *gin.Context)
	Delete(context *gin.Context)
	Update(context *gin.Context)
}

type userController struct {
	service *services.IUserService
}

func NewUserController(service *services.IUserService) *IUserController {

	var result IUserController = &userController{
		service: service,
	}

	return &result
}

func (userController *userController) Add(context *gin.Context) {
	var dto = UserDTO{}
	var err = context.BindJSON(&dto)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var service = *userController.service

	service.Add(services.User{
		Name: dto.Name,
		Age:  dto.Age,
	})
	context.String(http.StatusOK, `you was add a user called %s and %d years old`, dto.Name, dto.Age)
}

func (userController *userController) GetById(context *gin.Context) {

	var id, _ = context.Params.Get("id")
	var service = *userController.service

	var user = service.GetById(id)

	context.JSON(http.StatusOK, UserDTO{
		Name: user.Name,
		Age:  user.Age,
		Id:   user.Id,
	})

}

func (userController *userController) GetAll(context *gin.Context) {
	var service = *userController.service
	var users = service.GetAll()

	var result []UserDTO

	for _, user := range users {
		result = append(result, UserDTO{
			Name: user.Name,
			Age:  user.Age,
			Id:   user.Id,
		})
	}

	context.JSON(http.StatusOK, result)
}

func (userController *userController) Delete(context *gin.Context) {
	var id, _ = context.Params.Get("id")
	var service = *userController.service
	service.Delete(id)
}

func (userController *userController) Update(context *gin.Context) {
	var dto = UserDTO{}
	var err = context.ShouldBind(&dto)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var service = *userController.service
	service.Update(services.User{
		Name: dto.Name,
		Age:  dto.Age,
		Id:   dto.Id,
	})

}

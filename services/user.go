package services

import (
	"GoDEMO/daos"
	uuid "github.com/satori/go.uuid"
)

type (
	User struct {
		Id   string
		Name string
		Age  int
	}
	IUserService interface {
		Add(user User)
		GetById(id string) User
		GetAll() []User
		Delete(id string)
		Update(user User)
	}
	userService struct {
		userDAO *daos.IUserDAO
	}
)

func NewUserService(dao *daos.IUserDAO) *IUserService {

	var result IUserService = &userService{
		userDAO: dao,
	}

	return &result
}

func (service *userService) Add(user User) {
	var dao = *service.userDAO

	user.Id = uuid.NewV4().String()

	var po = daos.UserPO{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	}

	dao.Add(po)
}

func (service *userService) GetById(id string) User {
	var dao = *service.userDAO

	var po = dao.FindById(id)

	return User{
		Id:   po.Id,
		Name: po.Name,
		Age:  po.Age,
	}
}

func (service *userService) GetAll() []User {
	var dao = *service.userDAO

	var po = dao.FindAll()

	var result []User

	for _, current := range po {
		result = append(result, User{
			Id:   current.Id,
			Name: current.Name,
			Age:  current.Age,
		})
	}

	return result
}

func (service *userService) Delete(id string) {
	var dao = *service.userDAO
	dao.Delete(id)
}

func (service *userService) Update(user User) {
	var dao = *service.userDAO
	dao.Update(daos.UserPO{
		Id:   user.Id,
		Name: user.Name,
		Age:  user.Age,
	})
}

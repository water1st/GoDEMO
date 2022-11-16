package daos

type UserPO struct {
	Id   string
	Name string
	Age  int
}

type IUserDAO interface {
	Add(po UserPO)
	Update(po UserPO)
	Delete(id string)
	FindById(id string) UserPO
	FindAll() []UserPO
}

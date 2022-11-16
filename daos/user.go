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

type memoryUserDAO struct {
	users *[]UserPO
}

func NewMemoryUserDAO() *IUserDAO {
	var result IUserDAO = &memoryUserDAO{
		users: &[]UserPO{},
	}
	return &result
}

func (dao *memoryUserDAO) Add(po UserPO) {
	var users = *dao.users
	var newUsers = append(users, po)
	dao.users = &newUsers
}

func (dao *memoryUserDAO) Update(po UserPO) {
	var users = *dao.users

	for _, user := range users {
		if user.Id == po.Id {
			user.Name = po.Name
			user.Age = po.Age
		}
	}
}

func (dao *memoryUserDAO) Delete(id string) {
	var removeIndex int
	var users = *dao.users
	for index, user := range users {
		if user.Id == id {
			removeIndex = index
			break
		}
	}

	var newUsers = append(users[:removeIndex], users[removeIndex+1])
	dao.users = &newUsers
}

func (dao *memoryUserDAO) FindById(id string) UserPO {
	var users = *dao.users
	var result UserPO

	for _, user := range users {
		if user.Id == id {
			result = user
		}
	}

	return result
}

func (dao *memoryUserDAO) FindAll() []UserPO {
	return *dao.users
}

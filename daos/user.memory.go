package daos

type memoryUserDAO struct {
	users *[]UserPO
}

func newMemoryUserDAO() *IUserDAO {
	var result IUserDAO = &memoryUserDAO{
		users: &[]UserPO{},
	}
	return &result
}

func (memory *memoryUserDAO) Add(po UserPO) {
	var users = *memory.users
	var newUsers = append(users, po)
	memory.users = &newUsers
}

func (memory *memoryUserDAO) Update(po UserPO) {
	var users = *memory.users

	for _, user := range users {
		if user.Id == po.Id {
			user.Name = po.Name
			user.Age = po.Age
		}
	}
}

func (memory *memoryUserDAO) Delete(id string) {
	var removeIndex int
	var users = *memory.users
	for index, user := range users {
		if user.Id == id {
			removeIndex = index
			break
		}
	}

	var newUsers = append(users[:removeIndex], users[removeIndex+1])
	memory.users = &newUsers
}

func (memory *memoryUserDAO) FindById(id string) UserPO {
	var users = *memory.users
	var result UserPO

	for _, user := range users {
		if user.Id == id {
			result = user
		}
	}

	return result
}

func (memory *memoryUserDAO) FindAll() []UserPO {
	return *memory.users
}

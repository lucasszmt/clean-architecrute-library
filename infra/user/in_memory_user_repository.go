package user

import (
	"awesomeLibraryProject/domain/library/user"
)

type InMemoryUserRepository struct {
	Users []user.User
}

func (i *InMemoryUserRepository) FindById(id int) {
	panic("implement me")
}

func (i *InMemoryUserRepository) Delete(user *user.User) {
	panic("implement me")
}

func (i *InMemoryUserRepository) Update(user *user.User) {
	panic("implement me")
}

func (i *InMemoryUserRepository) Register(user *user.User) {
	i.Users = append(i.Users, *user)
}

//
//func (i *InMemoryUserRepository) Update(user *user.User) {
//	panic("implement me")
//}
//
//func (i *InMemoryUserRepository) FindById(id int) {
//	panic("implement me")
//}
//
//func (i *InMemoryUserRepository) Register(user user.User) {
//	i.Users = append(i.Users, user)
//}
//
//func (i *InMemoryUserRepository) Delete(user user.User) {
//	for index, elem := range i.Users {
//		if user.GetId() == elem.GetId() {
//			i.Users = append(i.Users[:index], i.Users[index+1:]...)
//		}
//	}
//}
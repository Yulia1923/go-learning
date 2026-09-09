package main

import "fmt"

type User struct {
	ID   int
	Name string
	Age  int
}

type UserService struct {
	users map[int]*User
}

func (s *UserService) AddUser(user *User) {
	s.users[user.ID] = user
}

func (s *UserService) UpdateUser(id int, name string, age int) {
	user, ok := s.users[id]

	if !ok {

		fmt.Println("该用户不存在")
		return
	}

	user.Name = name
	user.Age = age
}

func NewUserService() *UserService {
	return &UserService{
		users: make(map[int]*User),
	}
}
func main() {
	service := NewUserService()

	user := User{
		ID:   1,
		Name: "Tom",
		Age:  27,
	}
	service.AddUser(&user)

	fmt.Println("修改前:", user)

	service.UpdateUser(1, "Bob", 21)

	fmt.Println("修改后:", user)

}

package main

import (
	"errors"
	"fmt"
)

//User ...
type User struct {
	ID          int
	Username    string
	PhoneNumber string
}

//UserList ...
type UserList struct {
	Users []User
}

var (
	errUserNotFound     = errors.New("user not found")
	errUserExist        = errors.New("user exist")
	errIDExist          = errors.New("user With the given ID exist")
	errUsernameExist    = errors.New("user With the given username exist")
	errPhoneNumberExist = errors.New("user With the given phone number exist")
)

var users UserList

func main() {
	users.Create(User{
		ID:          1,
		Username:    "Dipto",
		PhoneNumber: "00000",
	})
	users.Create(User{
		ID:          2,
		Username:    "Mondal",
		PhoneNumber: "08088",
	})
	fmt.Println(users)
	users.Update(User{
		ID:          1,
		Username:    "dipto",
		PhoneNumber: "0100",
	})
	users.Update(User{
		ID:          2,
		Username:    "dipto",
		PhoneNumber: "0100",
	})
	fmt.Println(users)
	users.Search("0100")

}

//Create ...
func (users *UserList) Create(user User) error {
	idExist := users.validateID(user)
	phoneNumberExist := users.validatePhoneNumber(user)
	usernameExist := users.validateUserName(user)
	if !phoneNumberExist && !idExist && !usernameExist {
		users.Users = append(users.Users, user)
		return nil
	} else {
		return errUserExist
	}
}

//Update ...
func (users *UserList) Update(user User) error {
	phoneNumberExist := users.validatePhoneNumber(user)
	usernameExist := users.validateUserName(user)
	if !usernameExist {
		for k, v := range users.Users {
			if v.ID == user.ID {
				users.Users[k].Username = user.Username
				return nil
			}
		}
	} else if !phoneNumberExist {
		for k, v := range users.Users {
			if v.ID == user.ID {
				users.Users[k].PhoneNumber = user.PhoneNumber
				return nil
			}
		}
	}
	return errUserNotFound
}

//Fetch ...
func Fetch() {
	for _, v := range users.Users {
		fmt.Print(v.ID)
		fmt.Print("\t", v.Username)
		fmt.Println("\t", v.PhoneNumber)
	}
}

func (users *UserList) validateID(user User) bool {
	//fmt.Println(user)
	for _, v := range users.Users {
		if v.ID == user.ID {
			fmt.Println(errIDExist)
			return true
		}
	}
	return false
}

func (users *UserList) validateUserName(user User) bool {
	for _, v := range users.Users {
		if v.Username == user.Username {
			fmt.Println(errUsernameExist)
			return true
		}
	}
	return false
}

func (users *UserList) validatePhoneNumber(user User) bool {
	for _, v := range users.Users {
		if v.PhoneNumber == user.PhoneNumber {
			fmt.Println(errPhoneNumberExist)
			return true
		}
	}
	return false
}

//Search ...
func (users *UserList) Search(search string) error {
	for _, v := range users.Users {
		if v.PhoneNumber == search || v.Username == search {
			fmt.Println(v)
			return nil
		}
	}
	return errUserNotFound
}

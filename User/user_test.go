package main

import (
	"fmt"
	"log"
	"testing"
)

var userTest UserList

func TestUserList_Create(t *testing.T) {
	t.Run("Checking user creation", func(t *testing.T) {
		userTest.Create(User{
			ID:          1,
			Username:    "dipto",
			PhoneNumber: "01729677770",
		})

		notFound := true
		for _, v := range userTest.Users {
			if v.ID == 0 && v.Username == "dipto" && v.PhoneNumber == "01729677770" {
				log.Printf("expected %s, got %s", v.Username, v.Username)
				notFound = false
			}
		}
		if notFound {
			fmt.Errorf("expected %s, got nothing", "dipto")
		}
	})

	t.Run("Checking error for user creation", func(t *testing.T) {
		err := userTest.Create(User{
			ID:          0,
			Username:    "dipto",
			PhoneNumber: "011",
		})
		if err != nil {
			log.Printf("expected %s error, got %s error", errUserExist, err)
		} else {
			fmt.Errorf("expect %v got %v", errUserExist, err)
		}
	})
}

func TestUserList_Update(t *testing.T) {
	t.Run("Checking if a user is being updated", func(t *testing.T) {
		userTest.Create(User{
			ID:          3,
			Username:    "dipto01",
			PhoneNumber: "01729677771",
		})
		err := userTest.Update(User{
			ID:          3,
			Username:    "dipt01",
			PhoneNumber: "0172900000",
		})

		if err == nil {
			fmt.Println("expected nil, got nil")
		} else {
			fmt.Errorf("expected nil got %s", err)
		}
	})

	t.Run("Checking if a user is being updated if username or phone number exist in other entry", func(t *testing.T) {
		err := userTest.Update(User{
			ID:          3,
			Username:    "dipto",
			PhoneNumber: "0172900000",
		})

		if err != nil {
			log.Println(err)
		} else {
			fmt.Errorf("expected nil, got %s", err)
		}
	})

	t.Run("User Not Found", func(t *testing.T) {
		err := userTest.Update(User{
			ID:          0,
			Username:    "a",
			PhoneNumber: "01",
		})
		if err != nil{
			fmt.Printf("Expected nil error got nil errr")
		}else {
			fmt.Errorf("expected nil, got %s",err)
		}
	})
}

func TestUserList_Search(t *testing.T) {
	t.Run("Checking User is found", func(t *testing.T) {
		err := userTest.Search("dipto")
		if err != nil {
			fmt.Errorf("expected user with username dipto and phone number 01729677770, got %s", err)
		} else {
			log.Println("expected nil error, got nil error")
		}
	})

	t.Run("Checking if error is thrown when user is not found", func(t *testing.T) {
		err := userTest.Search("diptoNotFound")
		if err == nil {
			fmt.Errorf("expected %s, got %s", errUserNotFound, err)
		} else {
			log.Printf("expected %s, got %s", err, err)
		}
	})

}

func Test_validateId(t *testing.T) {
	t.Run("Checking if user with id exist", func(t *testing.T) {
		isExist := userTest.validateID(User{
			ID:          1,
			Username:    "abc",
			PhoneNumber: "01939999",
		})

		if isExist {
			log.Printf("expected true, got %t", isExist)
		} else {
			fmt.Errorf("expected false, got %t", isExist)
		}
	})

	t.Run("Checking if user not found with id", func(t *testing.T) {
		isExist := userTest.validateID(User{
			ID:          7,
			Username:    "abc",
			PhoneNumber: "01939999",
		})

		if !isExist {
			log.Printf("expected false, got %t", isExist)
		} else {
			fmt.Errorf("expected false, got %t", isExist)
		}
	})

}

func Test_validateUserName(t *testing.T) {
	t.Run("Checking if user with username exist", func(t *testing.T) {
		isExist := userTest.validateUserName(User{
			ID:          1,
			Username:    "dipto01",
			PhoneNumber: "0172900000",
		})

		if isExist {
			log.Printf("expected true, got %t", isExist)
		} else {
			fmt.Errorf("expected false, got %t", isExist)
		}
	})

	t.Run("Checking if user not found with username", func(t *testing.T) {
		isExist := userTest.validateUserName(User{
			ID:          7,
			Username:    "abcasd",
			PhoneNumber: "01939999",
		})

		if !isExist {
			log.Printf("expected false, got %t", isExist)
		} else {
			fmt.Errorf("expected false, got %t", isExist)
		}
	})
}

func Test_validatePhoneNumber(t *testing.T) {
	t.Run("Checking if user with phone number exist", func(t *testing.T) {
		isExist := userTest.validatePhoneNumber(User{
			ID:          1,
			Username:    "dipto01",
			PhoneNumber: "0172900000",
		})

		if isExist {
			log.Printf("expected true, got %t", isExist)
		} else {
			fmt.Errorf("expected false, got %t", isExist)
		}
	})

	t.Run("Checking if user not found with phone number", func(t *testing.T) {
		isExist := userTest.validatePhoneNumber(User{
			ID:          7,
			Username:    "abcasd",
			PhoneNumber: "01939999878878",
		})

		if !isExist {
			log.Printf("expected false, got %t", isExist)
		} else {
			fmt.Errorf("expected false, got %t", isExist)
		}
	})
}

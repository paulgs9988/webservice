package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)
//this returns a slice of pointers to User objects
func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New User must not include id or it must be set to 0")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

func getUserById(id int) (User, error) {
	//don't need index so use write-only variable (here we have index, user)
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
			return u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			//splicing syntax: take a slice of users up to but not including i
			// append to it all of the users from one past the match on
			// i.e.: [user1, user2, user3, user4, user5, user6]
			// index:   0      1      2      3      4      5
			// matching index is 1 so:
			// 0--> -- -->2-->3-->4-->5
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("User with ID '%v' not found", id)
}
package models

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	//a slice (i.e. dynamic array) that holds pointers to user objects
	//by working with pointers, we can manipulate users without having
	//to copy them, so it is more efficient
	users []*User
	nextID = 1
)
//this returns a slice of pointers to User objects
func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}
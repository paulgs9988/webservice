//User controller file
//This will handle hte routing of an HTTP request to the correct method that'll
//handle it
package controllers

import (
	"net/http"
	"regexp"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

//below we'll bind a function to a method (specify the type we'll bind to)
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GM From the Evacuation Zone!"))
}

//constructor
func newUserController() *userController {
	return &userController{
		//looking for "slash users slash" followed by number for extending 
		//user list in future
		userIDPattern: regexp.MustCompile(`^/users/(\d+)/?`),
	}
}
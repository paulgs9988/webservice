//User controller file
//This will handle hte routing of an HTTP request to the correct method that'll
//handle it
package controllers

import "net/http"

type userController struct {}

//below we'll bind a function to a method (specify the type we'll bind to)
func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from the User Controller!"))
}
package controller

import "fmt"

//@resp-custom-user
type DemoRest struct {
	Tel  int
	Time string
}

//@resp-custom-user
type UserRest struct {
	Tel  int
	Time string
}

// [name string, age int]
// @POST /registUser
func (receiver *UserRest) RegistUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return true
}

// [name string, age int]
// @GET /logOutUser
func (receiver *UserRest) LogOutUser(name string, age int) (success bool) {
	fmt.Println(name, age, "-----user")
	return false
}

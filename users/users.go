package users

import (
	"main/files"
	"main/plan"
	"strconv"
)

const PATH = "users"

func AddUser(user *User) {
	plan.CreatePlansList(user.ID)
	files.WriteFile(PATH, strconv.Itoa(user.ID), user)
}

func GetUSer(id int) User {
	user := &User{}
	files.ReadFile(user, PATH, strconv.Itoa(id))

	return *user
}

func ForEachUser(callbackFn func(user *User)) {
	//for _, value := range users {
	//	callbackFn(value)
	//}
}

package users

var users = map[int]*User{}

func AddUser(user *User) {
	users[user.ID] = user
}

func GetUSer(id int) *User {
	return users[id]
}

func ForEachUser(callbackFn func(user *User)) {
	for _, value := range users {
		callbackFn(value)
	}
}

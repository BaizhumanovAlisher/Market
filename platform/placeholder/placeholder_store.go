package placeholder

import (
	"platform/authorization/identity"
	"platform/services"
	"strings"
)

func RegisterPlaceholderUserStore() {
	err := services.AddSingleton(func() identity.UserStore {
		return &UserStore{}
	})

	if err != nil {
		panic(err)
	}
}

var users = map[int]identity.User{
	1: identity.NewBasicUser(1, "Alice", "Administrator"),
	2: identity.NewBasicUser(2, "Bob"),
}

type UserStore struct{}

func (store *UserStore) GetUserByID(id int) (identity.User, bool) {
	user, found := users[id]
	return user, found
}

func (store *UserStore) GetUserByName(name string) (identity.User, bool) {
	for _, user := range users {
		if strings.EqualFold(user.GetDisplayName(), name) {
			return user, true
		}
	}

	return nil, false
}

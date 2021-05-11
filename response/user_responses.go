package response

import "web/db/models"

type UserResource struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	City     string `json:"city"`
}

type UserCollection []UserResource

func NewUserResource(user models.User) UserResource {
	return UserResource{
		UserName: user.UserName,
		Name:     user.Name,
		Phone:    user.Phone,
		City:     user.City,
	}
}

func NewUserCollection(users []models.User) UserCollection {
	uCol := make(UserCollection, 0, len(users))
	for _, user := range users {
		ur := NewUserResource(user)
		uCol = append(uCol, ur)
	}
	return uCol
}

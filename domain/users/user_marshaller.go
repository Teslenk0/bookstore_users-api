package users

type PublicUser struct {
	ID          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	ID          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (user *User) Marshall(isPublic bool)  interface{}{
	if isPublic{
		return PublicUser{
			ID:          user.ID,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}else{
		return PrivateUser{
			ID:          user.ID,
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			DateCreated: user.DateCreated,
			Status:      user.Status,
		}
	}
}

func (users Users) Marshall(isPublic bool) []interface{}{
	result := make([]interface{}, len(users))

	for index, user := range users{
		result[index] = user.Marshall(isPublic)
	}
	return result
}
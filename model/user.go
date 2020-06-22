package model

// User structure
type User struct {
	ID        string `json:"id"`
	CreatedAt int64  `json:"createdAt"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UserEqual test the equality of the user field (except the id)
func UserEqual(expected, actual User) bool {
	if expected.Email != actual.Email {
		return false
	}
	if expected.FirstName != actual.FirstName {
		return false
	}
	if expected.LastName != actual.LastName {
		return false
	}
	return true
}

// CreateUserInput structure
type CreateUserInput struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// UpdateUserInput structure
type UpdateUserInput struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

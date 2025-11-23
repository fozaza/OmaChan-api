package table

// User is main data user
type User struct {
	Name  string
	Email string
	Level int
}

// UserInput for login
type UserInput struct {
	Email    string
	Password string
}

// UserReturn for jwt
type UserReturn struct {
	Email string
	Level int
}

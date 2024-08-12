package user

// CreateInput defines input to create a user.
type CreateInput struct {
	Name     string
	Email    string
	Password string
}

// AccessTokenCreateInput defines input to create a user access token.
type AccessTokenCreateInput struct {
	Email    string
	Password string
}

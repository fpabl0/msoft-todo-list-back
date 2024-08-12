package user

// User defines a User.
type User struct {
	ID           uint
	Name         string
	Email        string `gorm:"unique"`
	PasswordHash string
}

// AccessTokenData defines the data contained in a user access token.
type AccessTokenData struct {
	UserID   uint
	UserName string
}

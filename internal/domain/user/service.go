package user

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/msoft-g1/todo-list-backend/internal/errs"
	"golang.org/x/crypto/bcrypt"
)

const secretKey = "NfwOmo6iN1UsdZVUSGeiO2gCDsvbW19u9JcnUwPtcGKvclOz0FX7R3BKmq3hNW2lsmfr4G5zFdn4kRaIy7QcCZitZRDlH72L"

// Service defines user service.
type Service struct {
	repo Repository
}

// NewService creates a new service.
func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

// CreateUser creates a new user.
func (s *Service) CreateUser(input *CreateInput) (*User, error) {
	// TODO check min password len
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return s.repo.Create(&User{
		Name:         input.Name,
		Email:        input.Email,
		PasswordHash: string(hashBytes),
	})
}

// GetUserByID returns the user with the specified id.
func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

// GetAllUsers returns all the users.
func (s *Service) GetAllUsers() ([]*User, error) {
	return s.repo.FindAll()
}

// ===============================================================
// Auth related
// ===============================================================

// CreateAccessToken creates access token.
func (s *Service) CreateAccessToken(input *AccessTokenCreateInput) (string, error) {
	u, err := s.repo.FindbyEmail(input.Email)
	if err != nil {
		return "", err
	}
	if u == nil {
		return "", errs.New(errs.CodeNotFound, "No existe el usuario con el correo especificado")
	}
	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(input.Password))
	if err != nil {
		return "", errs.New(errs.CodeInvalidCreds, "Credenciales inválidas")
	}
	return s.makeAccessToken(&AccessTokenData{UserID: u.ID, UserName: u.Name})
}

// RenewAccessToken renews access token.
func (s *Service) RenewAccessToken(token string) (string, error) {
	data, err := s.ValidateAccessToken(token)
	if err != nil {
		return "", err
	}
	return s.makeAccessToken(data)
}

// ValidateAccessToken validates and parses a user access token.
func (s *Service) ValidateAccessToken(token string) (*AccessTokenData, error) {
	tk, err := jwt.Parse(token, func(tk *jwt.Token) (interface{}, error) {
		if _, ok := tk.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", tk.Header["alg"])
		}
		return []byte(secretKey), nil
	}, jwt.WithExpirationRequired())
	if err != nil {
		return nil, errs.NewWithError(errs.CodeInvalidToken, err)
	}
	claims, ok := tk.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errs.New(errs.CodeInvalidToken, "Claims no esperados")
	}
	var data AccessTokenData
	n, _ := strconv.Atoi(claims["sub"].(string))
	data.UserID = uint(n)
	data.UserName = claims["username"].(string)
	return &data, nil
}

// ===============================================================
// Utils
// ===============================================================

func (s *Service) makeAccessToken(data *AccessTokenData) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      fmt.Sprint(data.UserID),
		"username": data.UserName,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

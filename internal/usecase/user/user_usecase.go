package user

import (
	"errors"
	domain "splunk_soar_clone/internal/domain/entity"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetByUsername(username string) (*domain.User, error)
	CreateToken(token *domain.Token) error
	DeleteTokenByUserID(userID int64) error
	CreateUser(user *domain.User) error
}

type UserUseCase struct {
	userRepo UserRepository
	jwtKey   []byte
}

func NewUserUseCase(userRepo UserRepository, jwtKey []byte) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		jwtKey:   jwtKey,
	}
}

func (uc *UserUseCase) Login(username, password string) (*domain.Token, *domain.User, error) {
	user, err := uc.userRepo.GetByUsername(username)
	if err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	// Check password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return nil, nil, errors.New("invalid credentials")
	}

	// Generate tokens
	accessToken, err := uc.generateAccessToken(user)
	if err != nil {
		return nil, nil, err
	}

	refreshToken, err := uc.generateRefreshToken(user)
	if err != nil {
		return nil, nil, err
	}

	// Delete existing tokens
	if err := uc.userRepo.DeleteTokenByUserID(user.UserID); err != nil {
		return nil, nil, err
	}

	// Create new token record
	token := &domain.Token{
		UserID:       user.UserID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		Expiry:       time.Now().Add(24 * time.Hour),
	}

	if err := uc.userRepo.CreateToken(token); err != nil {
		return nil, nil, err
	}

	return token, user, nil
}

func (uc *UserUseCase) CreateUser(user *domain.User) (*domain.User, error) {
	// Check if user already exists
	if _, err := uc.userRepo.GetByUsername(user.Username); err == nil {
		return nil, errors.New("user already exists")
	}
	// check if email already exists
	if _, err := uc.userRepo.GetByUsername(user.Email); err == nil {
		return nil, errors.New("email already exists")
	}
	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hash)
	// Create user
	if err := uc.userRepo.CreateUser(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (uc *UserUseCase) generateAccessToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"role_id": user.RoleID,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(uc.jwtKey)
}

func (uc *UserUseCase) generateRefreshToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(uc.jwtKey)
}

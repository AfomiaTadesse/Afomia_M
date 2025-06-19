package usecase

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/AfomiaTadesse/Afomia_M/backend/domain"
	"github.com/AfomiaTadesse/Afomia_M/backend/repository"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase interface {
	Signup(user *domain.SignupRequest) (*domain.AuthResponse, error)
	Login(user *domain.LoginRequest) (*domain.AuthResponse, error)
}
type userUsecase struct {
	userRepo       repository.UserRepository
	jwtSecret      string
	contextTimeout time.Duration}

func NewUserUsecase(userRepo repository.UserRepository, jwtSecret string, timeout time.Duration) UserUsecase {
	return &userUsecase{
		userRepo:       userRepo,
		jwtSecret:      jwtSecret,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Signup(userReq *domain.SignupRequest) (*domain.AuthResponse, error) {
    // Validate input
    if err := validateSignupInput(userReq); err != nil {
        return &domain.AuthResponse{
            Success: false,
            Message: "Validation failed",
            Errors:  []string{err.Error()},
        }, nil
    }

    // Check if email or username already exists
    if _, err := uc.userRepo.FindByEmail(context.Background(), userReq.Email); err == nil {
        return &domain.AuthResponse{
            Success: false,
            Message: "Email already exists",
        }, nil
    }

    if _, err := uc.userRepo.FindByUsername(context.Background(), userReq.Username); err == nil {
        return &domain.AuthResponse{
            Success: false,
            Message: "Username already exists",
        }, nil
    }

    // Hash password
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userReq.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, err
    }

    // Create user
    user := &domain.User{
        Username: userReq.Username,
        Email:    userReq.Email,
        Password: string(hashedPassword),
    }

    if err := uc.userRepo.Create(context.Background(), user); err != nil {
        return nil, err
    }

    return &domain.AuthResponse{
        Success: true,
        Message: "User created successfully",
    }, nil
}

func validateSignupInput(user *domain.SignupRequest) error {
    // Email validation
    emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
    if !emailRegex.MatchString(user.Email) {
        return errors.New("invalid email format")
    }

    if matched, _ := regexp.MatchString(`^[a-zA-Z0-9]+$`, user.Username); !matched {
        return errors.New("username must be alphanumeric only")
    }

    if len(user.Password) < 8 {
        return errors.New("password must be at least 8 characters long")
    }
    if matched, _ := regexp.MatchString(`[A-Z]`, user.Password); !matched {
        return errors.New("password must contain at least one uppercase letter")
    }
    if matched, _ := regexp.MatchString(`[a-z]`, user.Password); !matched {
        return errors.New("password must contain at least one lowercase letter")
    }
    if matched, _ := regexp.MatchString(`[!@#$%^&*(),.?":{}|<>]`, user.Password); !matched {
        return errors.New("password must contain at least one special character")
    }

    return nil
}

func (uc *userUsecase) Login(req *domain.LoginRequest) (*domain.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), uc.contextTimeout)
	defer cancel()

	// Find user by email
	user, err := uc.userRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		return &domain.AuthResponse{
			Success: false,
			Message: "Invalid email or password",
		}, nil
	}

	// Compare password with hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &domain.AuthResponse{
			Success: false,
			Message: "Invalid email or password",
		}, nil
	}

	// Generate JWT token
	token, err := uc.generateJWTToken(user)
	if err != nil {
		return nil, err
	}

	return &domain.AuthResponse{
		Success: true,
		Message: "Login successful",
		Token:   token,
	}, nil
}

func (uc *userUsecase) generateJWTToken(user *domain.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  user.ID.Hex(),
		"email":    user.Email,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(uc.jwtSecret))
}
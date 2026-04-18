package services

import (
	"encoding/json"
	"errors"
	"hoodhire/config"
	repositories "hoodhire/internal/repositories"
	"hoodhire/structures/dto"
	models "hoodhire/structures/models"
	"hoodhire/utils"
	
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type AuthServices struct {
	Repo  *repositories.AuthRepo
	Redis *redis.Client
}

func NewAuthService(repository *repositories.AuthRepo) *AuthServices {
	return &AuthServices{Repo: repository}
}

type SignupTemp struct {
	Username string
	Email    string
	Password string
	Role     string
	Otp      string
}

func (s *AuthServices) SendOtp(SignupInput *models.User) (string, error) {
	if s.Repo.UserExist(SignupInput.Email) {
		return "", errors.New("user already exists")
	}
	hashed, err := utils.GeneratePassword(SignupInput.Password)
	if err != nil {
		return "", err
	}
	SignupInput.Password = hashed
	otp := utils.GenerateOTP()
	data := &SignupTemp{
		Username: SignupInput.Username,
		Email:    SignupInput.Email,
		Password: SignupInput.Password,
		Role:     SignupInput.Role,
		Otp:      otp,
	}

	verificationToken := uuid.New().String()
	jsonData, _ := json.Marshal(data)

	e := s.Redis.Set(
		config.Ctx,
		"verify:"+verificationToken,
		jsonData,
		5*time.Minute,
	).Err()
	if e != nil {
		return "", err
	}
	if err := utils.SendOTPmail(SignupInput.Email, otp); err != nil {
		return "", errors.New("failed to send mail")
	}
	return verificationToken, nil
}

func (s *AuthServices) Signup(token, otp string) (*models.User, error) {
	key := "verify:" + token
	data, err := s.Redis.Get(config.Ctx, key).Result()
	if err == redis.Nil {
		return nil, errors.New("verification expired or invalid")
	}
	if err != nil {
		return nil, err
	}
	var SignupData SignupTemp
	if err := json.Unmarshal([]byte(data), &SignupData); err != nil {
		return nil, err
	}
	if SignupData.Otp != otp {
		return nil, errors.New("invalid otp")
	}
	user := models.User{
		Email:    SignupData.Email,
		Username: SignupData.Username,
		Password: SignupData.Password,
		Role:     SignupData.Role,
	}
	if err := s.Repo.CreateUser(&user); err != nil {
		return nil, err
	}

	s.Redis.Del(config.Ctx, key)

	return &user, nil
}

func (s *AuthServices) ResendOTP(token string) error {
	key := "verify:" + token
	data, err := s.Redis.Get(config.Ctx, key).Result()

	if err == redis.Nil {
		return errors.New("no pending signup found or token expired")
	}
	if err != nil {
		return err
	}

	var signupData SignupTemp
	if err := json.Unmarshal([]byte(data), &signupData); err != nil {
		return err
	}

	newOTP := utils.GenerateOTP()
	signupData.Otp = newOTP

	jsonData, _ := json.Marshal(signupData)
	if err := s.Redis.Set(config.Ctx, key, jsonData, 5*time.Minute).Err(); err != nil {
		return err
	}

	if err := utils.SendOTPmail(signupData.Email, newOTP); err != nil {
		return errors.New("failed to send OTP email")
	}
	return nil
}

func (s *AuthServices) Login(email, password string) (*models.User, error) {
	if !s.Repo.UserExist(email) {
		return nil, errors.New("please create an account first")
	}
	user, err := s.Repo.GetUser(email)
	if err != nil {
		return nil, err
	}
	if !utils.ComparePass(user.Password, password) {
		return nil, errors.New("invalid email or password")
	}
	return user, nil
}


//~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~admin~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~


func (s *AuthServices) GetAllUsers() ([]models.User, error) {
    return s.Repo.GetallUsers()
}

func (s *AuthServices) GetUsers(role string, blocked *bool, page, limit int) ([]models.User, error) {
	offset := (page - 1) * limit
	return s.Repo.GetUsers(role, blocked, limit, offset)
}

func (s *AuthServices) GetUserByID(userID uint) (*models.User, error) {
    return s.Repo.GetUserByID(userID)
}

func (s *AuthServices) DeleteUser(userID uint) error {
    return s.Repo.DeleteUser(userID)
}

func (s *AuthServices) BlockUser(userID uint) error {
    return s.Repo.BlockUser(userID)
}

func (s *AuthServices) UnblockUser(userID uint) error {
    return s.Repo.UnblockUser(userID)
}

func (s *AuthServices) AdminCreateUser(input *dto.CreateUserDto) error {
    if s.Repo.UserExist(input.Email) {
        return errors.New("user already exists")
    }
    hashed, err := utils.GeneratePassword(input.Password)
    if err != nil {
        return err
    }
    user := &models.User{
        Username: input.Username,
        Email:    input.Email,
        Password: hashed,
        Role:     input.Role,
    }
    return s.Repo.CreateUser(user)
}
func (s *AuthServices) EditUser(userID uint, input *dto.CreateUserDto) error {
    u, err := s.Repo.GetUserByID(userID)
    if err != nil {
        return err
    }

    if input.Username != "" {
        u.Username = input.Username
    }
    if input.Email != "" {
        u.Email = input.Email
    }
    if input.Role != "" {
        u.Role = input.Role
    }
    if input.Password != "" {
        hashed, err := utils.GeneratePassword(input.Password)
        if err != nil {
            return err
        }
        u.Password = hashed
    }

    return s.Repo.UpdateUser(u)
}


func (s *AuthServices) ChangePassword(userID uint, input *dto.ChangePassword) error {
    user, err := s.Repo.GetUserByID(userID)
    if err != nil {
        return errors.New("user not found")
    }

    if !utils.ComparePass(user.Password, input.OldPassword) {
        return errors.New("old password is incorrect")
    }

    if input.NewPassword == input.OldPassword {
        return errors.New("new password must be different from old password")
    }

    hashed, err := utils.GeneratePassword(input.NewPassword)
    if err != nil {
        return err
    }

    return s.Repo.UpdatePassword(userID, hashed)
}
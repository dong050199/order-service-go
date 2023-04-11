package usecase

import (
	"context"
	"order-service/pkg/config"
	"order-service/pkg/constants"
	"order-service/pkg/infra"
	"order-service/service/model/entity"
	"order-service/service/model/request"
	"order-service/service/repository"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var timeNow = time.Now().UTC()

type IuserUsecase interface {
	UserRegister(ctx context.Context, req request.UserRegister) (token string, err error)
	UserLogin(ctx context.Context, req request.UserLogin) (token string, err error)
	DeleteUser(ctx context.Context, id uint, deletedBy string) (err error)
	UpdateUser(ctx context.Context, req request.UserRegister, updatedBy string) (err error)
}

type userUsecase struct {
	userRepo repository.IuserRepo
}

func NewUserUsecase(
	userRepo repository.IuserRepo,
) IuserUsecase {
	return &userUsecase{userRepo}
}

func (u *userUsecase) UserRegister(
	ctx context.Context, req request.UserRegister,
) (token string, err error) {
	// run in transaction
	tx, err := infra.BeginTransaction()
	if err != nil {
		return
	}

	defer infra.ReleaseTransaction(tx, err)

	userID, err := u.userRepo.CreateUser(tx, req.ToEntityModel())
	if err != nil {
		return
	}
	var userLogin string
	userLogin = req.Email
	if len(req.UserName) > 0 {
		userLogin = req.UserName
	}

	token, err = u.generateToken(userID, userLogin)
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) UserLogin(
	ctx context.Context,
	req request.UserLogin,
) (token string, err error) {
	// run in transaction
	tx, err := infra.BeginTransaction()
	if err != nil {
		return
	}

	defer infra.ReleaseTransaction(tx, err)

	valid, userID, err := u.userRepo.ValidateUser(entity.User{
		UserName: req.UserNameOrEmail,
		Email:    req.UserNameOrEmail,
		Password: req.Password,
	})

	if !valid {
		return
	}

	token, err = u.generateToken(userID, req.UserNameOrEmail)
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) DeleteUser(
	ctx context.Context,
	id uint,
	deletedBy string) (err error) {
	tx, err := infra.BeginTransaction()
	if err != nil {
		return
	}

	defer infra.ReleaseTransaction(tx, err)

	err = u.userRepo.DeleteUser(tx, entity.User{
		ID:        id,
		DeletedBy: deletedBy,
	})
	if err != nil {
		return
	}

	return
}

func (u *userUsecase) UpdateUser(
	ctx context.Context,
	req request.UserRegister,
	updatedBy string,
) (err error) {
	tx, err := infra.BeginTransaction()
	if err != nil {
		return
	}

	defer infra.ReleaseTransaction(tx, err)

	err = u.userRepo.UpdateUser(tx, entity.User{
		ID:        req.ID,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		UserName:  req.UserName,
		Email:     req.Email,
		Password:  req.Password,
		UpdatedAt: &timeNow,
		UpdatedBy: updatedBy,
	})
	if err != nil {
		return
	}

	return
}

func (v *userUsecase) generateToken(userID uint, userNameOrEmail string) (string, error) {
	lifespan, err := strconv.Atoi(config.JwtConfig().TokenHourLifeSpan)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{}
	claims[constants.Authorized] = true
	claims[constants.UserID] = userID
	claims[constants.Email] = userNameOrEmail
	claims[constants.ExpireDate] = time.Now().Add(time.Hour * time.Duration(lifespan)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.JwtConfig().APISecret))
}

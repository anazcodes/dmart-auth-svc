package usecase

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/repo/interfaces"
	services "github.com/anazibinurasheed/dmart-auth-svc/internal/usecase/interfaces"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/util"
)

const (
	email    = "email"
	phone    = "phone"
	username = "username"
)

var (
	ErrRecordAlreadyExist = errors.New("record already exist")
	ErrPasswordMismatch   = errors.New("password mismatch")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrNoAccount          = errors.New("no account record")
	ErrDuplicatingPhone   = errors.New("duplicating phone an account already registered with this phone")
	ErrDuplicatingEmail   = errors.New("duplicating email an account already registered with this email")
)

type userUseCase struct {
	UserRepo interfaces.UserRepo
}

func NewUserUseCase(userRepo interfaces.UserRepo) services.UserUseCase {
	return &userUseCase{
		UserRepo: userRepo,
	}
}

func (u *userUseCase) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) error {

	err := u.validateAccountRequest(ctx, req)
	if util.HasError(ctx, err) {
		return err
	}

	req.Password, err = util.HashPassword(req.Password)
	if util.HasError(ctx, err) {
		return err
	}

	t := payload.Time{
		Now: time.Now(),
	}

	err = u.UserRepo.CreateAccount(ctx, req, t)

	if util.HasError(ctx, err) {
		return err
	}

	return nil
}

// Check wether any user have already used this credentials to create-account
func (u *userUseCase) validateAccountRequest(ctx context.Context, req *pb.CreateAccountRequest) error {
	contact := payload.Contact{
		Email: req.Email,
		Phone: req.Phone,
	}

	account, err := u.UserRepo.GetMatchingAccountUsingPhone(ctx, contact)

	if util.HasError(ctx, err) {
		return err
	}

	if account.ID != 0 {
		return ErrDuplicatingPhone
	}

	account, err = u.UserRepo.GetMatchingAccountUsingEmail(ctx, contact)

	if util.HasError(ctx, err) {
		return err
	}

	if account.ID != 0 {
		return ErrDuplicatingEmail

	}

	return nil
}

func (u *userUseCase) UserLogin(ctx context.Context, req *pb.UserLoginRequest) error {
	method := util.GetLoginMethod(req)

	account, err := u.getAccountWithLoginCred(ctx, req, method)
	if util.HasError(ctx, err) {
		return err
	}

	if account.ID == 0 {
		return ErrNoAccount
	}

	if !util.CompareHashAndPassword(account.Password, req.Password) {
		return ErrPasswordMismatch
	}
	util.Logger("return reached")
	return nil
}

// It's a helper func of UserLogin it will find out the user account details from db.
// Mainly the function is for provide the better user experience by giving feature to login with any of the valid credentials used to create account.
func (u *userUseCase) getAccountWithLoginCred(ctx context.Context, req *pb.UserLoginRequest, method string) (payload.UserAccount, error) {
	util.Logger(fmt.Sprint(method + "---"))
	if method == email {
		account, err := u.UserRepo.GetMatchingAccountUsingEmail(ctx, payload.Contact{
			Email: req.LoginInput,
		})
		if util.HasError(ctx, err) {
			return payload.UserAccount{}, err
		}
		return account, nil
	}

	if method == phone {
		phne, err := strconv.Atoi(req.LoginInput)
		if util.HasError(ctx, err) {
			return payload.UserAccount{}, err
		}
		account, err := u.UserRepo.GetMatchingAccountUsingPhone(ctx, payload.Contact{
			Phone: int64(phne),
		})

		util.Logger(method, "phone", account)
		if util.HasError(ctx, err) {
			return payload.UserAccount{}, err
		}
		return account, nil
	}

	if method == username {
		account, err := u.UserRepo.GetUserAccountByName(ctx, req.LoginInput)
		if util.HasError(ctx, err) {
			return payload.UserAccount{}, err
		}
		return account, nil
	}
	return payload.UserAccount{}, nil
}

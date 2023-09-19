package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/repository/interfaces"
	services "github.com/anazibinurasheed/dmart-auth-svc/internal/usecase/interfaces"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/utils"
)

var (
	RecordAlreadyExist = errors.New("record already exist")
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

	if utils.HasError(ctx, err) {
		return err
	}

	t := payload.Time{
		Now: time.Now(),
	}

	err = u.UserRepo.CreateAccount(ctx, req, t)

	if utils.HasError(ctx, err) {
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

	if utils.HasError(ctx, err) {
		return err
	}

	if account.ID != 0 {
		return RecordAlreadyExist
	}

	account, err = u.UserRepo.GetMatchingAccountUsingEmail(ctx, contact)

	if utils.HasError(ctx, err) {
		return err
	}

	if account.ID != 0 {
		return RecordAlreadyExist

	}

	return nil
}

// func (s *Server) UserLogin(ctx context.Context, req pb.UserLoginRequest) {

// }

// func (s *Server) ValidateRegistration(ctx context.Context, req pb.UserLoginRequest) {

// }

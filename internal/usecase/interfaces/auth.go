package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

type UserUseCase interface {
	CreateAccount(context.Context, *pb.CreateAccountRequest) error
	UserLogin(context.Context, *pb.UserLoginRequest) (string, error)
	AdminLogin(context.Context, *pb.AdminLoginRequest) (string, error)
}

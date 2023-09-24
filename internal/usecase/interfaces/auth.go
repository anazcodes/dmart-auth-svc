package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

type UserUseCase interface {
	CreateAccount(context.Context, *pb.CreateAccountRequest) error
	UserLogin(context.Context, *pb.UserLoginRequest) (int64, error)
	AdminLogin(context.Context, *pb.AdminLoginRequest) error
	ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (int64, error)
}

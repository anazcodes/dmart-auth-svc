package usecase

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/config"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

func (u *userUseCase) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) error {
	authData := config.GetConfig()

	if req.Username != authData.Admin && req.Password != authData.AdminPassword {
		return ErrInvalidCredentials
	}

	return nil
}

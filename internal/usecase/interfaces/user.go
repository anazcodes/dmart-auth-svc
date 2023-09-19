package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

type UserUseCase interface {
	CreateAccount(context.Context, *pb.CreateAccountRequest) error
}

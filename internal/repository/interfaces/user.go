package interfaces

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

type UserRepo interface {
	CreateAccount(context.Context, *pb.CreateAccountRequest, payload.Time) error
	GetMatchingAccountUsingPhone(context.Context, payload.Contact) (payload.UserAccount, error)
	GetMatchingAccountUsingEmail(context.Context, payload.Contact) (payload.UserAccount, error)
}

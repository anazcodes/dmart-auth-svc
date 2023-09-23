package interfaces

import (
	contxt "context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

type context = contxt.Context

type UserRepo interface {
	CreateAccount(context, *pb.CreateAccountRequest, payload.Time) error
	GetMatchingAccountUsingPhone(context, payload.Contact) (payload.UserAccount, error)
	GetMatchingAccountUsingEmail(context, payload.Contact) (payload.UserAccount, error)
	GetUserAccountByID(context, uint) (payload.UserAccount, error)
	GetUserAccountByName(context, string) (payload.UserAccount, error)
}

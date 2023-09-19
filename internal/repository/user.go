package repository

import (
	contxt "context"
	"fmt"

	payload "github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/repository/interfaces"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/utils"
	"gorm.io/gorm"
)

// Type context.Context
type context = contxt.Context

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &userRepo{DB: db}
}

// func (u *userRepo) CreateAccount(ctx context, req *pb.CreateAccountRequest, t payload.Time) error {
// 	query := `insert into users (username,email,phone,password,created_at,updated_at)
// 		      values ($1,$2,$3,$4,$5,$6) returning *;`
// 	utils.HighlightError("create account")
// 	err := u.DB.Raw(query, req.Username, req.Email, req.Phone, req.Password, t.Now, t.Now).Error

// 	return err
// }

func (u *userRepo) CreateAccount(ctx context, req *pb.CreateAccountRequest, t payload.Time) error {
	var data payload.UserAccount
	query := `insert into users (username,email,phone,password,created_at,updated_at)
		      values ($1,$2,$3,$4,$5,$6) returning *;`
	utils.HighlightError("create account")
	err := u.DB.Raw(query, req.Username, req.Email, req.Phone, req.Password, t.Now, t.Now).Scan(&data).Error
	utils.HighlightError(fmt.Sprint(data))
	return err
}

func (u *userRepo) GetMatchingAccountUsingPhone(ctx context, req payload.Contact) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where phone = $1;`
	utils.HighlightError("GetMatchingAccountUsingPhone")
	err := u.DB.Raw(query, req.Phone).Scan(&data).Error
	utils.HighlightError(fmt.Sprint(data))
	utils.HighlightError(fmt.Sprint(req.Phone))

	return data, err
}

func (u *userRepo) GetMatchingAccountUsingEmail(ctx context, req payload.Contact) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where email = $1;`
	utils.HighlightError("GetMatchingAccountUsingEmail")

	err := u.DB.Raw(query, req.Email).Scan(&data).Error
	return data, err
}

func (u *userRepo) DropTable() {
	query := `drop table addresses ;drop table users;`
	utils.HighlightError("GetMatchingAccountUsingEmail")

	err := u.DB.Raw(query)
	utils.HighlightError(fmt.Sprint(err.Error))
	return
}

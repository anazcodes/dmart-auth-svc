package repo

import (
	contxt "context"
	"fmt"

	payload "github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/repo/interfaces"
	util "github.com/anazibinurasheed/dmart-auth-svc/internal/util"
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

func (u *userRepo) CreateAccount(ctx context, req *pb.CreateAccountRequest, t payload.Time) error {
	var data payload.UserAccount
	query := `insert into users (username,email,phone,password,created_at,updated_at)
		      values ($1,$2,$3,$4,$5,$6) returning *;`
	util.Logger("create account")
	err := u.DB.Raw(query, req.Username, req.Email, req.Phone, req.Password, t.Now, t.Now).Scan(&data).Error
	util.Logger(fmt.Sprint(data))
	return err
}

func (u *userRepo) GetMatchingAccountUsingPhone(ctx context, req payload.Contact) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where phone = $1;`
	util.Logger("GetMatchingAccountUsingPhone")
	err := u.DB.Raw(query, req.Phone).Scan(&data).Error
	util.Logger(fmt.Sprint(data))
	util.Logger(fmt.Sprint(req.Phone))

	return data, err
}

func (u *userRepo) GetMatchingAccountUsingEmail(ctx context, req payload.Contact) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where email = $1;`
	util.Logger("GetMatchingAccountUsingEmail")

	err := u.DB.Raw(query, req.Email).Scan(&data).Error
	return data, err
}

func (u *userRepo) GetUserAccountByName(ctx context, username string) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where username = $1;`
	err := u.DB.Raw(query, username).Scan(&data).Error
	return data, err
}

func (u *userRepo) GetUserAccountByID(ctx context, userID uint) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where id = $1;`
	err := u.DB.Raw(query, userID).Scan(&data).Error
	return data, err
}

func (u *userRepo) DropTable() {
	query := `drop table addresses ;drop table users;`
	util.Logger("GetMatchingAccountUsingEmail")

	err := u.DB.Raw(query)
	util.Logger(fmt.Sprint(err.Error))
}

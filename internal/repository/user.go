package repository

import (
	contxt "context"

	payload "github.com/anazibinurasheed/dmart-auth-svc/internal/payload"
	"gorm.io/gorm"
)

// Type context.Context
type context = contxt.Context

type userRepo struct {
	DB *gorm.DB
}

func (u *userRepo) CreateAccount(ctx context, req payload.UserAccount) error {
	query := `insert into users (username,email,phone,password,created_at,updated_at)
		      values ($1,$2,$3,$4,$5,$6);`

	err := u.DB.Raw(query, req.Username, req.Email, req.Phone, req.Password, req.CreatedAt, req.UpdatedAt).Error
	return err
}

func (u *userRepo) GetAccountWithLoginData(ctx context, req payload.UserLogin) (payload.UserAccount, error) {
	var data payload.UserAccount
	query := `select * from users where username = $1 or phone = $2 or email = $3;`
	err := u.DB.Raw(query, req.LoginInput, req.LoginInput, req.LoginInput).Scan(&data).Error
	return data, err
}

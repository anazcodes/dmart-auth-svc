package di

import (
	"github.com/anazibinurasheed/dmart-auth-svc/internal/config"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/db"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/repo"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/services"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/usecase"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/util"
)

func InitializeDeps(cfg config.Config) (*services.Server, error) {
	db, err := db.InitDB(cfg)
	if util.HasError(err) {
		return &services.Server{}, err
	}

	userRepo := repo.NewUserRepo(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	return &services.Server{
		UserUseCase: userUseCase,
	}, nil

}

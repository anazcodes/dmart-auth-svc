package di

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/config"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/db"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/repository"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/services"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/usecase"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/utils"
)

func InitialiazeDeps(cfg config.Config) (*services.Server, error) {
	db, err := db.InitDB(cfg)
	if utils.HasError(context.Background(), err) {
		return &services.Server{}, err
	}

	userRepo := repository.NewUserRepo(db)
	userUseCase := usecase.NewUserUseCase(userRepo)

	return &services.Server{
		UserUseCase: userUseCase,
	}, nil

}

package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/usecase"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/usecase/interfaces"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/utils"
)

type Server struct {
	UserUseCase interfaces.UserUseCase
	pb.UnimplementedAuthServiceServer
}

func (s *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	fmt.Println("i am reached here")
	err := s.UserUseCase.CreateAccount(ctx, req)
	utils.HighlightError(fmt.Sprint(req.Phone))
	if err == usecase.RecordAlreadyExist {
		return &pb.CreateAccountResponse{
			Status: http.StatusConflict,
			Msg:    "user already exist with this credentials",
			Error:  err.Error(),
		}, nil

	}

	if utils.HasError(ctx, err) {
		return &pb.CreateAccountResponse{
			Status: http.StatusInternalServerError,
			Msg:    "failed",
			Error:  err.Error(),
		}, nil
	}

	return &pb.CreateAccountResponse{
		Status: http.StatusCreated,
		Msg:    "success, account created",
	}, nil

}

func (s *Server) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {
	return &pb.AdminLoginResponse{}, nil
}

func (s *Server) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return &pb.UserLoginResponse{}, nil
}

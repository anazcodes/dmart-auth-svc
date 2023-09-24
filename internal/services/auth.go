package services

import (
	"context"
	"net/http"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/usecase"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/usecase/interfaces"
	"github.com/anazibinurasheed/dmart-auth-svc/internal/util"
)

type Server struct {
	UserUseCase interfaces.UserUseCase
	pb.UnimplementedAuthServiceServer
}

func (s *Server) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {

	err := s.UserUseCase.CreateAccount(ctx, req)
	if err == usecase.ErrDuplicatingEmail {
		return &pb.CreateAccountResponse{
			Status: http.StatusConflict,
			Msg:    "user already have an account with this email",
			Error:  err.Error(),
		}, nil

	}

	if err == usecase.ErrDuplicatingPhone {
		return &pb.CreateAccountResponse{
			Status: http.StatusConflict,
			Msg:    "already have an account with this phone",
			Error:  err.Error(),
		}, nil

	}

	if util.HasError(err) {
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

func (s *Server) UserLogin(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	userID, err := s.UserUseCase.UserLogin(ctx, req)
	if err == usecase.ErrNoAccount {
		return &pb.UserLoginResponse{
			Status: http.StatusNotAcceptable,
			Msg:    "user don't have an account",
			Error:  err.Error(),
		}, nil
	}

	if err == usecase.ErrPasswordMismatch {
		return &pb.UserLoginResponse{
			Status: http.StatusNotAcceptable,
			Msg:    "incorrect password",
			Error:  err.Error(),
		}, nil

	}

	var role = "user"
	token, err := util.GenerateToken(uint(userID), role)

	if util.HasError(err) {
		return &pb.UserLoginResponse{
			Status: http.StatusInternalServerError,
			Msg:    "failed to generate token",
			Error:  err.Error(),
		}, nil
	}

	return &pb.UserLoginResponse{
		Status: http.StatusAccepted,
		Msg:    "login success",
		Token:  token,
	}, nil
}

func (s *Server) AdminLogin(ctx context.Context, req *pb.AdminLoginRequest) (*pb.AdminLoginResponse, error) {

	err := s.UserUseCase.AdminLogin(ctx, req)

	if util.HasError(err) {
		return &pb.AdminLoginResponse{
			Status: http.StatusNotAcceptable,
			Msg:    "invalid credentials",
			Error:  err.Error(),
		}, nil
	}

	var role = "suAdmin"
	token, err := util.GenerateToken(0, role)

	if util.HasError(err) {
		return &pb.AdminLoginResponse{
			Status: http.StatusInternalServerError,
			Msg:    "failed to generate token",
			Error:  err.Error(),
		}, nil
	}

	return &pb.AdminLoginResponse{
		Status: http.StatusOK,
		Msg:    "login success",
		Token:  token,
	}, nil
}

func (s *Server) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.ValidateTokenResponse, error) {

	userID, err := s.UserUseCase.ValidateToken(ctx, req)
	if util.HasError(err) {
		return &pb.ValidateTokenResponse{
			Status: http.StatusUnauthorized,
		}, nil
	}
	return &pb.ValidateTokenResponse{
		Status: http.StatusOK,
		UserID: userID,
	}, nil
}

package services

import (
	"context"

	"github.com/anazibinurasheed/dmart-auth-svc/internal/pb"
)

type Server struct {
}

func (s *Server) CreateAccount(ctx context.Context, req pb.CreateAccountRequest) {

}

func (s *Server) AdminLogin(ctx context.Context, req pb.AdminLoginRequest) {

}

func (s *Server) UserLogin(ctx context.Context, req pb.UserLoginRequest) {

}

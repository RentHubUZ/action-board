package service

import (
	pb "action-board/genproto/request"
	"action-board/internal/storage"
	"action-board/internal/storage/postgres"
	"context"
	"database/sql"
	"log/slog"
)

type RequestService struct {
	Request storage.IStorage
	Log     *slog.Logger
	pb.UnimplementedRequestServiceServer
}

func NewRequestService(db *sql.DB, log *slog.Logger) *RequestService {
	return &RequestService{
		Request: postgres.NewPostgresStorage(db, log),
		Log:     log,
	}
}

func (s *RequestService) CreateRequest(ctx context.Context, req *pb.CreateRequestRequest) (*pb.CreateRequestResponse, error) {
	s.Log.Info("Create request", slog.Any("request", req))
	res, err := s.Request.Request().Create(ctx, req)
	if err != nil {
		s.Log.Error("failed to create request", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

func (s *RequestService) GetRequest(ctx context.Context, req *pb.GetRequestRequest) (*pb.GetRequestResponse, error) {
	s.Log.Info("Get request", slog.Any("request", req))
	res, err := s.Request.Request().Get(ctx, req)
	if err != nil {
		s.Log.Error("failed to get request", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

func (s *RequestService) DeleteRequest(ctx context.Context, req *pb.DeleteRequestRequest) (*pb.Void, error) {
	s.Log.Info("Delete request", slog.Any("request", req))
	res, err := s.Request.Request().Delete(ctx, req)
	if err != nil {
		s.Log.Error("failed to delete request", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

package service

import (
	pb "action-board/genproto/notification"
	"action-board/internal/storage"
	"action-board/internal/storage/postgres"
	"context"
	"database/sql"
	"log/slog"
)

type NotificationService struct {
	Notification storage.IStorage
	Log          *slog.Logger
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(db *sql.DB, log *slog.Logger) *NotificationService {
	return &NotificationService{
		Notification: postgres.NewPostgresStorage(db, log),
		Log:          log,
	}
}

func (s *NotificationService) Create(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	res, err := s.Notification.Notification().Create(ctx, req)

	if err != nil {
		s.Log.Error("failed to create notification", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

func (s *NotificationService) Get(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {

	res, err := s.Notification.Notification().Get(ctx, req)
	if err != nil {
		s.Log.Error("failed to get notification", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

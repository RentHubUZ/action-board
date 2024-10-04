package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	accopb "action-board/genproto/accommodation"
	pb "action-board/genproto/notification"
	"action-board/internal/config"
	client "action-board/internal/pkg"
	logger "action-board/internal/pkg/logs"
	"action-board/internal/storage"

	"github.com/google/uuid"
)

type NotificationRepository struct {
	Db                   *sql.DB
	Log                  *slog.Logger
	AccommodationService accopb.AccommodationServiceClient
}

func NewNotificationRepository(db *sql.DB) storage.INotificationStorage {
	cfg := config.Load()
	return &NotificationRepository{
		Db:                   db,
		Log:                  logger.NewLogger(),
		AccommodationService: client.NewAccommodationClient(&cfg),
	}
}

// CreateNotification funksiyasi
func (r *NotificationRepository) Create(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.CreateNotificationResponse, error) {
	tx, err := r.Db.BeginTx(ctx, nil)
	if err != nil {
		r.Log.Error("failed to start transaction", slog.Any("error", err))
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback() 
		} else {
			err = tx.Commit() 
		}
	}()

	query := `INSERT INTO notifications (id, user_id, content, is_read) VALUES ($1, $2, $3, $4)`
	id := uuid.New().String()

	_, err = tx.ExecContext(ctx, query, id, req.UserId, req.Content, req.IsRead)
	if err != nil {
		r.Log.Error("failed to create notification", slog.Any("error", err))
		return nil, err
	}

	notification := &pb.Notification{
		Id:        id,
		UserId:    req.UserId,
		Content:   req.Content,
		IsRead:    req.IsRead,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	return &pb.CreateNotificationResponse{Notification: notification}, nil
}

// GetNotification funksiyasi
func (r *NotificationRepository) Get(ctx context.Context, req *pb.GetNotificationRequest) (*pb.GetNotificationResponse, error) {
	query := `SELECT id, user_id, content, is_read, created_at FROM notifications WHERE id = $1`

	var notification pb.Notification
	err := r.Db.QueryRowContext(ctx, query, req.Id).Scan(
		&notification.Id,
		&notification.UserId,
		&notification.Content,
		&notification.IsRead,
		&notification.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			r.Log.Warn("notification not found", slog.String("id", req.Id))
			return nil, fmt.Errorf("notification not found")
		}
		r.Log.Error("failed to get notification", slog.Any("error", err))
		return nil, err
	}

	return &pb.GetNotificationResponse{Notification: &notification}, nil
}

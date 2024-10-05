package postgres

import (
	"context"
	"database/sql"
	"log/slog"

	accopb "action-board/genproto/accommodation"
	pb "action-board/genproto/request"
	"action-board/internal/config"
	client "action-board/internal/pkg"
	logger "action-board/internal/pkg/logs"
	"action-board/internal/storage"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

type RequestRepository struct {
	Db                   *sql.DB
	Log                  *slog.Logger
	AccommodationService accopb.AccommodationServiceClient
}

func NewRequestRepository(db *sql.DB) storage.IRequestStorage {
	cfg := config.Load()
	return &RequestRepository{
		Db:                   db,
		Log:                  logger.NewLogger(),
		AccommodationService: client.NewAccommodationClient(&cfg),
	}
}

func (r *RequestRepository) Create(ctx context.Context, req *pb.CreateRequestRequest) (*pb.CreateRequestResponse, error) {
	query := `INSERT INTO requests (id, user_id, email, roommate_count, phone_number) 
	          VALUES ($1, $2, $3, $4, $5)`

	id := uuid.New().String()

	_, err := r.Db.ExecContext(ctx, query, id, req.UserId, pq.Array(req.Email), req.RoommateCount, req.PhoneNumber)
	if err != nil {
		r.Log.Error("failed to create request", slog.Any("error", err))
		return nil, err
	}

	request := &pb.Request{
		Id:            id,
		UserId:        req.UserId,
		Email:         req.Email,
		RoommateCount: req.RoommateCount,
		PhoneNumber:   req.PhoneNumber,
	}

	return &pb.CreateRequestResponse{Request: request}, nil
}

func (r *RequestRepository) Get(ctx context.Context, req *pb.GetRequestRequest) (*pb.GetRequestResponse, error) {
	query := `SELECT id, user_id, email, roommate_count, phone_number FROM requests WHERE id = $1`

	var request pb.Request

	err := r.Db.QueryRowContext(ctx, query, req.Id).Scan(
		&request.Id,
		&request.UserId,
		&request.Email,
		&request.RoommateCount,
		&request.PhoneNumber,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			r.Log.Warn("request not found", slog.String("id", req.Id))
			return nil, nil
		}
		r.Log.Error("failed to get request", slog.Any("error", err))
		return nil, err
	}

	return &pb.GetRequestResponse{Request: &request}, nil
}

func (r *RequestRepository) Delete(ctx context.Context, req *pb.DeleteRequestRequest) (*pb.Void, error) {
	query := `DELETE FROM requests WHERE id = $1`

	_, err := r.Db.ExecContext(ctx, query, req.Id)
	if err != nil {
		r.Log.Error("failed to delete request", slog.Any("error", err))
		return nil, err
	}

	r.Log.Info("request deleted successfully", slog.String("id", req.Id))
	return &pb.Void{}, nil
}
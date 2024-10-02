package postgres

import (
	accopb "action-board/genproto/accommodation"
	pb "action-board/genproto/report"
	"action-board/internal/config"
	client "action-board/internal/pkg"
	logger "action-board/internal/pkg/logs"
	"action-board/internal/storage"
	"context"
	"database/sql"
	"log/slog"
	"time"
)

type ReportRepository struct {
	Db                   *sql.DB
	Log                  *slog.Logger
	AccommodationService accopb.AccommodationServiceClient
}

func NewReportRepository(db *sql.DB) storage.IReportStorage {
	cfg := config.Load()
	return &ReportRepository{
		Db:                   db,
		Log:                  logger.NewLogger(),
		AccommodationService: client.NewAccommodationClient(&cfg),
	}
}

func (r *ReportRepository) Create(ctx context.Context, req *pb.CreateReportRequest) (*pb.CreateReportResponse, error) {
	query := `INSERT INTO reports (property_id, user_id, email, problem, description) 
	          VALUES ($1, $2, $3, $4, $5)`

	_, err := r.Db.ExecContext(ctx, query, req.PropertyId, req.UserId, req.Email, req.Problem, req.Description)
	if err != nil {
		r.Log.Error("failed to create report", slog.Any("error", err))
		return nil, err
	}

	report := &pb.Report{
		PropertyId:  req.PropertyId,
		UserId:      req.UserId,
		Email:       req.Email,
		Problem:     req.Problem,
		Description: req.Description,
		CreatedAt:   time.Now().Format("2006-01-02 15:04:05"),
	}

	return &pb.CreateReportResponse{Report: report}, nil
}

func (r *ReportRepository) Get(ctx context.Context, req *pb.GetReportRequest) (*pb.GetReportResponse, error) {
	query := `SELECT id, property_id, user_id, email, problem, description, created_at FROM reports WHERE id = $1`

	row := r.Db.QueryRowContext(ctx, query, req.Id)

	var report pb.Report
	err := row.Scan(&report.Id, &report.PropertyId, &report.UserId, &report.Email, &report.Problem, &report.Description, &report.CreatedAt)
	if err != nil {
		r.Log.Error("failed to get report", slog.Any("error", err))
		return nil, err
	}

	return &pb.GetReportResponse{Report: &report}, nil
}

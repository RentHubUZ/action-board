package service

import (
	pbf "action-board/genproto/report"
	"action-board/internal/storage"
	"action-board/internal/storage/postgres"
	"context"
	"database/sql"
	"log/slog"
)

type ReportService struct {
	pbf.UnimplementedReportServiceServer
	Report storage.IStorage
	Log    *slog.Logger
}

func NewReportService(db *sql.DB, log *slog.Logger) *ReportService {
	return &ReportService{
		Report: postgres.NewPostgresStorage(db, log),
		Log:    log,
	}
}

func (s *ReportService) Create(ctx context.Context, req *pbf.CreateReportRequest) (*pbf.CreateReportResponse, error) {
	s.Log.Info("Create report", slog.String("email", req.Email), slog.Any("request", req))
	res, err := s.Report.Report().Create(ctx, req)

	if err != nil {
		s.Log.Error("failed to create report", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

func (s *ReportService) Get(ctx context.Context, req *pbf.GetReportRequest) (*pbf.GetReportResponse, error) {
	s.Log.Info("Get report", slog.Any("request", req))
	res, err := s.Report.Report().Get(ctx, req)
	if err != nil {
		s.Log.Error("failed to get report", slog.Any("error", err))
		return nil, err
	}
	return res, nil
}

package service

import (
	pbr "action_board/genproto/reviews"
	"action_board/storage"
	"action_board/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type ReviewService struct {
	pbr.UnimplementedReviewsServer
	Review storage.IStorage
	Log    *slog.Logger
}

func NewReviewService(db *sql.DB, log *slog.Logger) *ReviewService {
	return &ReviewService{
		Review: postgres.NewPostgresStorage(db, log),
		Log:    log,
	}
}

func (r *ReviewService) CreateReview(ctx context.Context, req *pbr.CreateReviewReq) (*pbr.CreateReviewRes, error) {
	resp,err := r.Review.Review().CreateReview(ctx,req)
	if err != nil {
		r.Log.Error(fmt.Sprintf("Review ni create qilishda service qismida xatolik: %v",err.Error()))
		return nil,err
	}
	return resp,nil
}

func (r *ReviewService) GetAllReview(ctx context.Context, req *pbr.GetAllReviewReq) (*pbr.GetAllReviewRes, error) {
	resp,err:=r.Review.Review().GetAllReview(ctx,req)
	if err!= nil {
		r.Log.Error(fmt.Sprintf("Review ni get all qilishda xatolik: %v",err.Error()))
		return nil,err
	}
	return resp,nil
}

func (r *ReviewService) GetByIdReview(ctx context.Context, req *pbr.GetByIdReviewReq) (*pbr.GetByIdReviewRes, error) {
	resp,err:=r.Review.Review().GetByIdReview(ctx,req)
	if err!=nil{
		r.Log.Error(fmt.Sprintf("Review ni id buyicha uqib olishda xatolik: %v",err.Error()))
		return nil,err
	}
	return resp,nil
}

func (r *ReviewService) DeleteReview(ctx context.Context, req *pbr.DeleteReviewReq) (*pbr.DeleteReviewRes, error) {
	resp,err:=r.Review.Review().DeleteReview(ctx,req)
	if err!=nil{
		r.Log.Error(fmt.Sprintf("Review ni delete qilishda xatolik: %v",err.Error()))
		return nil,err
	}
	return resp,nil
}

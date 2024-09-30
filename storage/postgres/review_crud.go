package postgres

import (
	"action_board/config"
	pb "action_board/genproto/reviews"
	"action_board/pkg/logger"
	"action_board/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	accopb "action_board/genproto/accommodation"
	client "action_board/pkg"

	"github.com/google/uuid"
)

type ReviewsRepository struct {
	Db                   *sql.DB
	Log                  *slog.Logger
	AccommodationService accopb.AccommodationServiceClient
}

func NewReviewsRepository(db *sql.DB) storage.IReviewsStorage {
	cfg := config.Load()
	return &ReviewsRepository{
		Db:                   db,
		Log:                  logger.NewLogger(),
		AccommodationService: client.NewAccommodationClient(&cfg),
	}
}

func (r *ReviewsRepository) CreateReview(ctx context.Context, req *pb.CreateReviewReq) (*pb.CreateReviewRes, error) {
	req_accommodation := accopb.GetByIdHouseReq{
		Id: req.PropertyId,
	}

	res_accommodation, err := r.AccommodationService.GetByIdHouse(ctx, &req_accommodation)
	if err != nil {
		r.Log.Error(fmt.Sprintf("error when referring to getbyid in another service: %v", err.Error()))
		return nil, err
	}

	if res_accommodation.Id == "" {
		r.Log.Error(fmt.Sprintf("Property with ID %s not found", req.PropertyId))
		return nil, fmt.Errorf("property with ID %s not found", req.PropertyId)
	}

	query_review := `insert into reviews (
						id, user_id, property_id, rating, comment, created_at, updated_at
					) values (
					 	$1, $2, $3, $4, $5, $6, $7)`

	review_id := uuid.NewString()
	newtime := time.Now()

	_, err = r.Db.ExecContext(ctx, query_review, review_id, req.UserId, req.PropertyId, req.Rating, req.Comment, newtime, newtime)
	if err != nil {
		r.Log.Error(fmt.Sprintf("Error generating review data: %v", err.Error()))
		return nil, err
	}

	return &pb.CreateReviewRes{
		Id:         review_id,
		UserId:     req.UserId,
		PropertyId: req.PropertyId,
		Rating:     req.Rating,
		Comment:    req.Comment,
		CreatedAt:  newtime.Format(time.RFC3339),
		UpdatedAt:  newtime.Format(time.RFC3339),
	}, nil
}

func (r *ReviewsRepository) GetAllReview(ctx context.Context, req *pb.GetAllReviewReq) (*pb.GetAllReviewRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `select 
				id,	user_id, property_id, rating, comment, created_at, updated_at
			from 
				reviews
			where 
				deleted_at is null
			order by created_at desc
			LIMIT $1 OFFSET $2`

	rows, err := r.Db.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		r.Log.Error(fmt.Sprintf("Error getting all review data: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var reviews []*pb.Review

	for rows.Next() {
		var review pb.Review
		err := rows.Scan(
			&review.Id,
			&review.UserId,
			&review.PropertyId,
			&review.Rating,
			&review.Comment,
			&review.CreatedAt,
			&review.UpdatedAt,
		)
		if err != nil {
			r.Log.Error(fmt.Sprintf("error when reading all the information of the review: %v", err.Error()))
			return nil, err
		}
		reviews = append(reviews, &review)
	}

	if err := rows.Err(); err != nil {
		r.Log.Error(fmt.Sprintf("error review: %v", err.Error()))
		return nil, err
	}

	return &pb.GetAllReviewRes{
		Review: reviews,
	}, nil
}

func (r *ReviewsRepository) GetByIdReview(ctx context.Context, req *pb.GetByIdReviewReq) (*pb.GetByIdReviewRes, error) {
	query := `select 
				id,	user_id, property_id, rating, comment, created_at, updated_at
			from 
				reviews
			where 
				id = $1 and deleted_at is null`

	review := pb.GetByIdReviewRes{}
	err := r.Db.QueryRowContext(ctx, query, req.Id).Scan(&review.Id, &review.UserId,
		&review.PropertyId, &review.Rating, &review.Comment, &review.CreatedAt, &review.UpdatedAt)
	if err != nil {
		r.Log.Error(fmt.Sprintf("Error retrieving information about reiew's id: %v", err.Error()))
		return nil,err
	}

	return &review, nil
}

func (r *ReviewsRepository) DeleteReview(ctx context.Context, req *pb.DeleteReviewReq) (*pb.DeleteReviewRes, error) {
	query := `update 
				reviews
			set
				deleted_at = $1
			where 
				id = $2 and deleted_at is null`

	_, err := r.Db.ExecContext(ctx, query, req.Id)
	if err != nil {
		r.Log.Error(fmt.Sprintf("Error deleting reference by id: %v",err.Error()))
		return nil,err
	}
	return &pb.DeleteReviewRes{
		Message: "Your comment has been successfully deleted",
	},nil
}

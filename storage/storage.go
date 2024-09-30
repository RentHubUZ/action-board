package storage

import (
	pb "action_board/genproto/reviews"
	pbf "action_board/genproto/favorites"
	"context"
)

type IStorage interface {
	Review() IReviewsStorage
	Favorites() IFavoritesStorage
	Close()
}

type IReviewsStorage interface {
	CreateReview(ctx context.Context, req *pb.CreateReviewReq) (*pb.CreateReviewRes, error)
	GetAllReview(ctx context.Context, req *pb.GetAllReviewReq) (*pb.GetAllReviewRes, error)
	GetByIdReview(ctx context.Context, req *pb.GetByIdReviewReq) (*pb.GetByIdReviewRes, error)
	DeleteReview(ctx context.Context, req *pb.DeleteReviewReq) (*pb.DeleteReviewRes, error)
}

type IFavoritesStorage interface {
	CreateFavorites(ctx context.Context, req *pbf.CreateFavoritesReq) (*pbf.CreateFavoritesRes, error)
	GetAllFavorites(ctx context.Context, req *pbf.GetAllFavoritesReq) (*pbf.GetAllFavoritesRes, error)
	GetByIdFavorites(ctx context.Context, req *pbf.GetByIdFavoritesReq) (*pbf.GetByIdFavoritesRes, error)
	DeleteFavorites(ctx context.Context, req *pbf.DeleteFavoritesReq) (*pbf.DeleteFavoritesRes, error)
}
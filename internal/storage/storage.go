package storage

import (
	pbf "action-board/genproto/favorites"
	pb "action-board/genproto/reviews"
	pbr "action-board/genproto/report"
	pbrq "action-board/genproto/request"
	pbnt "action-board/genproto/notification"
	"context"
)

type IStorage interface {
	Review() IReviewsStorage
	Favorites() IFavoritesStorage
	Report() IReportStorage
	Request() IRequestStorage
	Notification() INotificationStorage
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

type IReportStorage interface {
	Create(ctx context.Context, req *pbr.CreateReportRequest) (*pbr.CreateReportResponse, error)
	Get(ctx context.Context, req *pbr.GetReportRequest) (*pbr.GetReportResponse, error)
}

type IRequestStorage interface {
	Create(ctx context.Context, req *pbrq.CreateRequestRequest) (*pbrq.CreateRequestResponse, error)
	Get(ctx context.Context, req *pbrq.GetRequestRequest) (*pbrq.GetRequestResponse, error)
}

type INotificationStorage interface {
	Create(ctx context.Context, req *pbnt.CreateNotificationRequest) (*pbnt.CreateNotificationResponse, error)
	Get(ctx context.Context, req *pbnt.GetNotificationRequest) (*pbnt.GetNotificationResponse, error)
}
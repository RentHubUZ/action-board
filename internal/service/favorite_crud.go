package service

import (
	pbf "action-board/genproto/favorites"
	"action-board/internal/storage"
	"action-board/internal/storage/postgres"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

type FavoriteService struct {
	pbf.UnimplementedFavoritesServer
	Favorite storage.IStorage
	Log      *slog.Logger
}

func NewFavoriteService(db *sql.DB, log *slog.Logger) *FavoriteService {
	return &FavoriteService{
		Favorite: postgres.NewPostgresStorage(db, log),
		Log:      log,
	}
}

func (f *FavoriteService) CreateFavorites(ctx context.Context, req *pbf.CreateFavoritesReq) (*pbf.CreateFavoritesRes, error) {
	resp, err := f.Favorite.Favorites().CreateFavorites(ctx, req)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Favorite ni create qilishda service qismida xatolik: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (f *FavoriteService) GetAllFavorites(ctx context.Context, req *pbf.GetAllFavoritesReq) (*pbf.GetAllFavoritesRes, error) {
	resp, err := f.Favorite.Favorites().GetAllFavorites(ctx, req)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Favorite ni barcha malumotlarini olishda service qismida xatolik: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (f *FavoriteService) GetByIdFavorites(ctx context.Context, req *pbf.GetByIdFavoritesReq) (*pbf.GetByIdFavoritesRes, error) {
	resp, err := f.Favorite.Favorites().GetByIdFavorites(ctx, req)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Favorite ni id buyicha malumotlarni olishda service qismidav xatolik: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (f *FavoriteService) DeleteFavorites(ctx context.Context, req *pbf.DeleteFavoritesReq) (*pbf.DeleteFavoritesRes, error) {
	resp, err := f.Favorite.Favorites().DeleteFavorites(ctx, req)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Favorite ni delete qilishda service qismida xatolik: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

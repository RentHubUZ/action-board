package postgres

import (
	"action_board/config"
	accopb "action_board/genproto/accommodation"
	pbf "action_board/genproto/favorites"
	client "action_board/pkg"
	"action_board/pkg/logger"
	"action_board/storage"
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type FavoriteRepository struct {
	Db                   *sql.DB
	Log                  *slog.Logger
	AccommodationService accopb.AccommodationServiceClient
}

func NewFavoritesRepository(db *sql.DB) storage.IFavoritesStorage {
	cfg := config.Load()
	return &FavoriteRepository{
		Db:                   db,
		Log:                  logger.NewLogger(),
		AccommodationService: client.NewAccommodationClient(&cfg),
	}
}

func (f *FavoriteRepository) CreateFavorites(ctx context.Context, req *pbf.CreateFavoritesReq) (*pbf.CreateFavoritesRes, error) {
	req_accommodation := accopb.GetByIdHouseReq{
		Id: req.PropertyId,
	}

	res_accommodation, err := f.AccommodationService.GetByIdHouse(ctx, &req_accommodation)
	if err != nil {
		f.Log.Error(fmt.Sprintf("error when referring to getbyid in another service: %v", err.Error()))
		return nil, err
	}

	if res_accommodation.Id == "" {
		f.Log.Error(fmt.Sprintf("Property with ID %s not found", req.PropertyId))
		return nil, fmt.Errorf("property with ID %s not found", req.PropertyId)
	}

	query := `insert into favorites (
				id, user_id, property_id, created_at
			) values (
			 	$1, $2, $3, $4)`

	favorite_id := uuid.NewString()
	newtime := time.Now()

	_, err = f.Db.ExecContext(ctx, query, favorite_id, req.UserId, req.PropertyId, newtime)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Error saving favorite: %v", err.Error()))
		return nil, err
	}

	return &pbf.CreateFavoritesRes{
		Id:         favorite_id,
		UserId:     req.UserId,
		PropertyId: req.PropertyId,
		CreatedAt:  newtime.Format(time.RFC3339),
	}, nil
}

func (f *FavoriteRepository) GetAllFavorites(ctx context.Context, req *pbf.GetAllFavoritesReq) (*pbf.GetAllFavoritesRes, error) {
	offset := (req.Page - 1) * req.Limit

	query := `select 
				id, user_id, property_id, created_at
			from 
				favorites
			order by created_at desc
			LIMIT $1 OFFSET $2`

	rows, err := f.Db.QueryContext(ctx, query, req.Limit, offset)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Error getting all favorite data: %v", err.Error()))
		return nil, err
	}
	defer rows.Close()

	var favorites []*pbf.Favorite
	for rows.Next() {
		var favorite pbf.Favorite
		err := rows.Scan(
			&favorite.Id,
			&favorite.UserId,
			&favorite.PropertyId,
			&favorite.CreatedAt,
		)
		if err != nil {
			f.Log.Error(fmt.Sprintf("error when reading all the information of the favorite: %v", err.Error()))
			return nil, err
		}

		favorites = append(favorites, &favorite)
	}

	if err := rows.Err(); err != nil {
		f.Log.Error(fmt.Sprintf("error favorite: %v", err.Error()))
		return nil, err
	}

	return &pbf.GetAllFavoritesRes{
		Favorites: favorites,
	},nil
}

func (f *FavoriteRepository) GetByIdFavorites(ctx context.Context, req *pbf.GetByIdFavoritesReq) (*pbf.GetByIdFavoritesRes, error) {
	query := `select 
				id, user_id, property_id, created_at
			from 
				favorites
			where
				id = $1`

	favorite := pbf.GetByIdFavoritesRes{}
	err := f.Db.QueryRowContext(ctx,query,req.Id).Scan(
				&favorite.Id,
				&favorite.UserId,
				&favorite.PropertyId,
				&favorite.CreatedAt,
	)	
	if err != nil {
		f.Log.Error(fmt.Sprintf("Error retrieving information about favorite's id: %v", err.Error()))
		return nil,err
	}
	return &favorite,nil
}

func (f *FavoriteRepository) DeleteFavorites(ctx context.Context, req *pbf.DeleteFavoritesReq) (*pbf.DeleteFavoritesRes, error) {
	query := `delete from favorites where id = $1`

	_,err := f.Db.ExecContext(ctx,query,req.Id)
	if err != nil {
		f.Log.Error(fmt.Sprintf("Error deleting reference by id: %v",err.Error()))
		return nil,err
	}

	return &pbf.DeleteFavoritesRes{
		Message: "Your comment has been successfully deleted",
	},nil
}

package main

import (
	"action_board/config"
	"action_board/pkg/logger"
	"action_board/service"
	"action_board/storage/postgres"
	"log"
	"net"
	pbf "action_board/genproto/favorites"
	pbr "action_board/genproto/reviews"

	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", config.Load().ACTION_BOARD)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	db, err := postgres.ConnectionDb()
	if err != nil {
		log.Fatal(err)
	}

	logs := logger.NewLogger()
	serviceReview := service.NewReviewService(db,logs)
	serviceFavorite := service.NewFavoriteService(db,logs)

	server := grpc.NewServer()
	pbr.RegisterReviewsServer(server,serviceReview)
	pbf.RegisterFavoritesServer(server,serviceFavorite)
	

	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

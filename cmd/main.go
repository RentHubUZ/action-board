package main

import (
	pbf "action-board/genproto/favorites"
	pbr "action-board/genproto/reviews"
	pbrq "action-board/genproto/request"
	pbnt "action-board/genproto/notification"
	pbrep "action-board/genproto/report"
	"action-board/internal/config"
	logger "action-board/internal/pkg/logs"
	"action-board/internal/storage/postgres"
	"action-board/internal/service"
	"log"
	"net"

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
	serviceReview := service.NewReviewService(db, logs)
	serviceFavorite := service.NewFavoriteService(db, logs)
	serviceReport := service.NewReportService(db, logs)
	serviceRequest := service.NewRequestService(db, logs)
	serviceNotification := service.NewNotificationService(db, logs)

	server := grpc.NewServer()
	pbr.RegisterReviewsServer(server, serviceReview)
	pbf.RegisterFavoritesServer(server, serviceFavorite)
	pbrep.RegisterReportServiceServer(server, serviceReport)
	pbrq.RegisterRequestServiceServer(server, serviceRequest)
	pbnt.RegisterNotificationServiceServer(server, serviceNotification)

	log.Printf("Server is listening on port %s\n", config.Load().ACTION_BOARD)
	err = server.Serve(listener)
	if err != nil {
		log.Fatal(err)
	}
}

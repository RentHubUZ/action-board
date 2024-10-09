package client

import (
	"action-board/internal/config"
	accopb "action-board/genproto/accommodation"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccommodationClient(cfg *config.Config) accopb.AccommodationServiceClient {
	conn, err := grpc.Dial("accommadation"+cfg.ACCOMMODATION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error while connecting authentication service: %v", err)
	}

	return accopb.NewAccommodationServiceClient(conn)
}

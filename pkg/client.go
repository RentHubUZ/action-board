package client

import (
	"log"
	"action_board/config"
	accopb "action_board/genproto/accommodation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAccommodationClient(cfg *config.Config) accopb.AccommodationServiceClient {
	conn, err := grpc.Dial("localhost" + cfg.ACCOMMODATION_SERVICE, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("error while connecting authentication service: %v", err)
	}

	return accopb.NewAccommodationServiceClient(conn)
}

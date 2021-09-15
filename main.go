package main

import (
	"PositionService/internal/config"
	protocol "PositionService/internal/protocol"
	"PositionService/internal/server"
	"PositionService/internal/service"
	"context"
	"fmt"
	"github.com/INEFFABLE-games/PriceService/models"
	protocol2 "github.com/INEFFABLE-games/PriceService/protocol"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

func main() {
	cfg := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// dial server
	conn, err := grpc.Dial(fmt.Sprintf(":%s", cfg.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	//start position grpc server
	go func() {
		grpcServer := grpc.NewServer()
		positionServer := server.NewPositionServer()
		protocol.RegisterPositionServiceServer(grpcServer,positionServer)

		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", cfg.PosGrpcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()

	// create stream
	clientFromService := protocol2.NewPriceServiceClient(conn)

	stream, err := clientFromService.Send(ctx)
	if err != nil {
		log.Fatalf("unable to open stream %v", err)
	}

	priceChannel := make(chan []models.Price)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	priceService := service.NewPriceService(priceChannel)

	//starts listening stream and refresh current price list
	go func() {
		priceService.Refresh(ctx, stream)
	}()

	<-c
	cancel()
	os.Exit(1)
}

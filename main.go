package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/INEFFABLE-games/PositionService/internal/config"
	"github.com/INEFFABLE-games/PositionService/internal/repository"
	"github.com/INEFFABLE-games/PositionService/internal/server"
	"github.com/INEFFABLE-games/PositionService/internal/service"
	"github.com/INEFFABLE-games/PositionService/protocol"
	"github.com/INEFFABLE-games/PriceService/models"
	protocol2 "github.com/INEFFABLE-games/PriceService/protocol"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
)

var currentPrices = []models.Price{}
var c = make(chan os.Signal, 1)

func main() {

	cfg := config.NewConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// dial server
	conn, err := grpc.Dial(fmt.Sprintf(":%s", cfg.GrpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("can not connect with server %v", err)
	}

	sqlConn, err := sql.Open("postgres", cfg.PsqlURI)
	if err != nil {
		log.Errorf("unable to connect with postgres %v", err)
	}

	positionRepository := repository.NewPositionRepository(sqlConn)
	positionService := service.NewPositionService(positionRepository, &currentPrices)

	//start position grpc server
	go func() {
		grpcServer := grpc.NewServer()
		positionServer := server.NewPositionServer(positionService)
		protocol.RegisterPositionServiceServer(grpcServer, positionServer)

		lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%v", cfg.PosGrpcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		err = grpcServer.Serve(lis)
		if err != nil {
			log.Errorf("unable to start grpc server %v", err.Error())
		}
	}()

	// create stream
	clientFromService := protocol2.NewPriceServiceClient(conn)

	stream, err := clientFromService.Send(ctx)
	if err != nil {
		log.Fatalf("unable to open stream %v", err)
	}

	signal.Notify(c, os.Interrupt)

	priceService := service.NewPriceService(&currentPrices)

	//starts listening stream and refresh current price list
	go func() {
		priceService.Refresh(ctx, stream)
	}()

	<-c
	cancel()
	os.Exit(1)
}

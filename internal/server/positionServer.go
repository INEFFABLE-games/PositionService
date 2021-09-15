package server

import (
	"PositionService/internal/protocol"
	"context"
	log "github.com/sirupsen/logrus"
)

type PositionServer struct {
	protocol.UnimplementedPositionServiceServer
}

func (p *PositionServer) Buy(ctx context.Context, request *protocol.BuyRequest)(*protocol.BuyReply,error){

	log.Info(request.UserToken)
	log.Info(request.Price)

	return nil,nil
}

func (p *PositionServer) Sell(ctx context.Context, request *protocol.SellRequest)(*protocol.SellReply,error){
	return nil,nil
}

func (p *PositionServer) Get(ctx context.Context, request *protocol.GetRequest)(*protocol.GetReply,error){
	return nil,nil
}


func NewPositionServer()*PositionServer{
	return &PositionServer{}
}

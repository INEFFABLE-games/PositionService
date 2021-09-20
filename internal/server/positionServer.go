package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/INEFFABLE-games/PositionService/internal/service"
	"github.com/INEFFABLE-games/PositionService/protocol"
	"github.com/INEFFABLE-games/PriceService/models"
	log "github.com/sirupsen/logrus"
)

type PositionServer struct {
	positionService *service.PositionService

	protocol.UnimplementedPositionServiceServer
}

func (p *PositionServer) Buy(ctx context.Context, request *protocol.BuyRequest) (*protocol.BuyReply, error) {

	price := models.Price{}
	owner := request.GetUserToken()

	err := json.Unmarshal(request.GetPrice(), &price)
	if err != nil {
		mes := fmt.Sprintf("unable to buy price: %v", err.Error())
		return &protocol.BuyReply{Message: &mes}, err
	}

	err = p.positionService.Buy(ctx, price, owner)
	if err != nil {
		log.WithFields(log.Fields{
			"handler ": "positionService",
		}).Errorf("unable to insert price into db %v", err.Error())
		mes := err.Error()

		return &protocol.BuyReply{Message: &mes}, err
	}

	mes := fmt.Sprintf("succes buyed position: %v for user: %v", price, request.GetUserToken())
	return &protocol.BuyReply{Message: &mes}, err
}

func (p *PositionServer) Sell(ctx context.Context, request *protocol.SellRequest) (*protocol.SellReply, error) {

	price := models.Price{}
	userToken := request.GetUserToken()

	err := json.Unmarshal(request.GetPrice(), &price)
	if err != nil {
		mes := fmt.Sprintf("unable to buy price: %v", err.Error())
		return &protocol.SellReply{Message: &mes}, err
	}

	err = p.positionService.Sell(ctx, price, userToken)

	mes := fmt.Sprintf("succes selled position: %v for user: %v", price, request.GetUserToken())
	return &protocol.SellReply{Message: &mes}, err
}

func (p *PositionServer) Get(ctx context.Context, request *protocol.GetRequest) (*protocol.GetReply, error) {
	userToken := request.GetUserToken()

	res, err := p.positionService.GetByOwner(ctx, userToken)
	if err != nil {
		return &protocol.GetReply{}, err
	}

	marshalRes, err := json.Marshal(res)
	if err != nil {
		return &protocol.GetReply{}, err
	}

	return &protocol.GetReply{ButchOfPrices: marshalRes}, err
}

func NewPositionServer(positionService *service.PositionService) *PositionServer {
	return &PositionServer{positionService: positionService}
}

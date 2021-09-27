package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/INEFFABLE-games/PositionService/internal/service"
	"github.com/INEFFABLE-games/PositionService/protocol"
	"github.com/INEFFABLE-games/PriceService/models"
	authService "github.com/INEFFABLE-games/authService/models"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

// PositionServer structure for PositionServer object
type PositionServer struct {
	positionService *service.PositionService

	protocol.UnimplementedPositionServiceServer
}

func middlewareJWTValidation(token string) (string, error) {
	parsedToken, err := jwt.ParseWithClaims(
		token,
		&authService.CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte("operationToken"), nil
		},
	)
	if err != nil {
		return "", errors.New("token expired")
	}

	claim, ok := parsedToken.Claims.(*authService.CustomClaims)
	if !ok {
		return "", err
	}

	return claim.Uid, err
}

// Buy rpc buy function
func (p *PositionServer) Buy(ctx context.Context, request *protocol.BuyRequest) (*protocol.BuyReply, error) {

	price := models.Price{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Errorf("unable to get metadata")
	}

	token := md.Get("Token")[0]

	uid, err := middlewareJWTValidation(token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler ": "positionServer",
		}).Errorf("unable to validate token %v", err.Error())
		mes := err.Error()

		return &protocol.BuyReply{Message: &mes}, err
	}

	err = json.Unmarshal(request.GetPrice(), &price)
	if err != nil {
		mes := fmt.Sprintf("unable to buy price: %v", err.Error())
		return &protocol.BuyReply{Message: &mes}, err
	}

	err = p.positionService.Buy(ctx, price, uid)
	if err != nil {
		log.WithFields(log.Fields{
			"handler ": "positionServer",
		}).Errorf("unable to insert price into db %v", err.Error())
		mes := err.Error()

		return &protocol.BuyReply{Message: &mes}, err
	}

	mes := fmt.Sprintf("succes buyed position: %v for user: %v", price, token)
	return &protocol.BuyReply{Message: &mes}, err
}

// Sell rpc sell function
func (p *PositionServer) Sell(ctx context.Context, request *protocol.SellRequest) (*protocol.SellReply, error) {

	price := models.Price{}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Errorf("unable to get metadata")
	}

	token := md.Get("Token")[0]

	uid, err := middlewareJWTValidation(token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler ": "positionServer",
		}).Errorf("unable to validate token %v", err.Error())
		mes := err.Error()

		return &protocol.SellReply{Message: &mes}, err
	}

	err = json.Unmarshal(request.GetPrice(), &price)
	if err != nil {
		mes := fmt.Sprintf("unable to buy price: %v", err.Error())
		return &protocol.SellReply{Message: &mes}, err
	}

	err = p.positionService.Sell(ctx, price, uid)

	mes := fmt.Sprintf("succes selled position: %v for user: %v", price, token)
	return &protocol.SellReply{Message: &mes}, err
}

// Get rpc get function
func (p *PositionServer) Get(ctx context.Context, request *protocol.GetRequest) (*protocol.GetReply, error) {
	//token := request.GetUserToken()

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Errorf("unable to get metadata")
	}

	token := md.Get("Token")[1]

	uid, err := middlewareJWTValidation(token)
	if err != nil {
		log.WithFields(log.Fields{
			"handler ": "positionServer",
		}).Errorf("unable to validate token %v", err.Error())

		return &protocol.GetReply{}, err
	}

	res, err := p.positionService.GetByOwner(ctx, uid)
	if err != nil {
		return &protocol.GetReply{}, err
	}

	marshalRes, err := json.Marshal(res)
	if err != nil {
		return &protocol.GetReply{}, err
	}

	return &protocol.GetReply{ButchOfPrices: marshalRes}, err
}

// NewPositionServer creates PositionServer object
func NewPositionServer(positionService *service.PositionService) *PositionServer {
	return &PositionServer{positionService: positionService}
}

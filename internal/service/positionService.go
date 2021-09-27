package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/INEFFABLE-games/PositionService/internal/repository"
	"github.com/INEFFABLE-games/PriceService/models"
	"github.com/INEFFABLE-games/PriceService/protocol"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
	"io"
)

// PositionService structure for PositionService object
type PositionService struct {
	currentPrices []models.Price
	positionRepo  *repository.PositionRepository
	pricesForPNL  map[string]map[string]chan models.Price
}

// GetPNL calculate and returns pnl for position
func (p *PositionService) GetPNL(lastPrice models.Price, pos models.Price) int64 {
	var pnl int64
	if lastPrice.Name == pos.Name {
		pnl = int64(lastPrice.Bid - pos.Ask)
	}
	return pnl
}

// isFresh check is position up to date
func (p *PositionService) isFresh(price models.Price) bool {

	for _, v := range p.currentPrices {
		if v.Name == price.Name {
			log.Info(v)
			log.Info(price)
		}
		if v.Id == price.Id {
			return true
		}
	}

	return false
}

// Buy process buy request
func (p *PositionService) Buy(ctx context.Context, price models.Price, owner string) error {
	if p.isFresh(price) {

		currentBalance, err := p.positionRepo.GetBalance(ctx, owner)
		if err != nil {
			return err
		}

		currentBalance -= int64(price.Ask)

		err = p.positionRepo.UpdateBalance(ctx, owner, currentBalance)
		if err != nil {
			return err
		}

		return p.positionRepo.Insert(ctx, price, owner)
	}
	return errors.New("price isn't fresh")
}

// Refresh get fresh prices from grpc stream and write into currentPrices map
func (p *PositionService) Refresh(ctx context.Context, stream protocol.PriceService_SendClient) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			data, err := stream.Recv()
			if err == io.EOF {
				log.WithFields(log.Fields{
					"handler": "main",
					"action":  "get data from stream",
				}).Errorf("End of file %v", err.Error())
			}

			butchOfPrices := []models.Price{}

			err = json.Unmarshal(data.ButchOfPrices, &butchOfPrices)
			if err != nil {
				log.WithFields(log.Fields{
					"handler": "main",
					"action":  "unmarshal butch of prices",
				}).Errorf("unable to unmarshal butch of prices %v", err.Error())
			}
			go func() {
				p.currentPrices = butchOfPrices
				//log.Infof(fmt.Sprintf("[%v] NPL =  %v", v, p.positionService.GetPNL(v)))

				log.Infof("current prices listenning  %v", p.pricesForPNL)

				for _, v := range butchOfPrices {
					if p.pricesForPNL[v.Name] != nil {
						for _, v1 := range p.pricesForPNL[v.Name] {
							v1 <- v
						}
					}
				}
			}()
		}
	}
}

// Sell process sell request
func (p *PositionService) Sell(ctx context.Context, price models.Price, owner string) error {
	if p.isFresh(price) {

		var storedPrice models.Price
		found := false

		allUserPrices, err := p.positionRepo.GetByOwner(ctx, owner)
		if err != nil {
			return err
		}

		for _, v := range allUserPrices {
			if v.Name == price.Name && v.Ask == 0 {
				storedPrice = v
				found = true
			}
		}
		if !found {
			return errors.New("unable to find price")
		}

		currentBalance, err := p.positionRepo.GetBalance(ctx, owner)
		if err != nil {
			return err
		}

		currentBalance += int64(p.GetPNL(price, storedPrice))

		err = p.positionRepo.UpdateBalance(ctx, owner, currentBalance)
		if err != nil {
			return err
		}

		return p.positionRepo.Update(ctx, price, owner)
	}
	return errors.New("price isn't fresh")
}

// GetByOwner call getByOwner function on repository
func (p *PositionService) GetByOwner(ctx context.Context, owner string) ([]models.Price, error) {
	return p.positionRepo.GetByOwner(ctx, owner)
}

// ListenNotify starts pg notify listener
func (p *PositionService) ListenNotify(ctx context.Context) error {
	p.pricesForPNL = make(map[string]map[string]chan models.Price)
	_, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "positions",
			User:     "postgres",
			Password: "12345",
		},
		AfterConnect: func(conn *pgx.Conn) error {
			err := conn.Listen("db_notifications_open")
			if err != nil {
				return err
			}

			err = conn.Listen("db_notifications_close")
			if err != nil {
				return err
			}

			for {
				msg, err := conn.WaitForNotification(ctx)
				if err != nil {
					return err
				}

				pos := models.Price{}
				if err = json.Unmarshal([]byte(msg.Payload), &pos); err != nil {
					log.Errorf("error while unmarchaling: %v", err)
				}

				if msg.Channel == "db_notifications_open" {

					if p.pricesForPNL[pos.Name] == nil {
						p.pricesForPNL[pos.Name] = make(map[string]chan models.Price)
					}
					p.pricesForPNL[pos.Name][pos.Id] = make(chan models.Price)

					go func() {
						for {
							select {
							case lastPrice, find := <-p.pricesForPNL[pos.Name][pos.Id]:
								if !find {
									log.Info("can't find current price...")
									return
								}

								log.Infof(fmt.Sprintf("[%v] PNL =  %v", pos.Name, p.GetPNL(lastPrice, pos)))
							}
						}
					}()
				}
				if msg.Channel == "db_notifications_close" {

					if ch, ok := p.pricesForPNL[pos.Name][pos.Id]; ok {
						close(ch)
					}
					delete(p.pricesForPNL[pos.Name], pos.Id)
				}

				continue
			}
		},
	})
	if err != nil {
		return err
	}
	return nil
}

// NewPositionService creates PositionService object
func NewPositionService(ctx context.Context, positionRepo *repository.PositionRepository, currentPrices []models.Price, pricesForPNL map[string]map[string]chan models.Price) *PositionService {

	service := &PositionService{
		positionRepo:  positionRepo,
		currentPrices: currentPrices,
		pricesForPNL:  pricesForPNL,
	}

	go func() {
		err := service.ListenNotify(ctx)
		if err != nil {
			log.Errorf("unable to start listenning pg notify")
		}
	}()

	return service
}

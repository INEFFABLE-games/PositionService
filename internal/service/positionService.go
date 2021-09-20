package service

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/INEFFABLE-games/PositionService/internal/repository"
	"github.com/INEFFABLE-games/PriceService/models"
	"github.com/jackc/pgx"
	log "github.com/sirupsen/logrus"
	"time"
)

type PositionService struct {
	currentPrices *[]models.Price
	positionRepo  *repository.PositionRepository
}

func (p *PositionService) getPNL(price models.Price) int {
	var pnl int
	for _, v := range *p.currentPrices {
		if v.Name == price.Name {
			pnl = int(price.Ask - v.Bid)
		}
	}
	return pnl
}

func (p *PositionService) isFresh(price models.Price) bool {

	for _, v := range *p.currentPrices {
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

func (p *PositionService) Buy(ctx context.Context, price models.Price, owner string) error {
	if p.isFresh(price) {
		return p.positionRepo.Insert(ctx, price, owner)
	}
	return errors.New("price isn't fresh")
}

func (p *PositionService) Sell(ctx context.Context, price models.Price, owner string) error {
	if p.isFresh(price) {

		found := false

		allUserPrices, err := p.positionRepo.GetByOwner(ctx, owner)
		if err != nil {
			return err
		}

		for _, v := range allUserPrices {
			if v.Name == price.Name {
				found = true
			}
		}
		if !found {
			return errors.New("unable to find price")
		}

		return p.positionRepo.Update(ctx, price, owner)
	}
	return errors.New("price isn't fresh")
}

func (p *PositionService) GetByOwner(ctx context.Context, owner string) ([]models.Price, error) {
	return p.positionRepo.GetByOwner(ctx, owner)
}

func (p *PositionService) ListenNotify() error {
	_, err := pgx.NewConnPool(pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:     "localhost",
			Port:     5432,
			Database: "positions",
			User:     "postgres",
			Password: "12345",
		},
		AfterConnect: func(conn *pgx.Conn) error {
			err := conn.Listen("db_notifications")
			if err != nil {
				return err
			}

			for {
				msg, err := conn.WaitForNotification(context.Background())
				if err != nil {
					return err
				}

				pos := models.Price{}

				if msg.Channel == "db_notifications" {
					go func() {
						ctx, cancel := context.WithCancel(context.Background())
						defer cancel()
						ticker := time.NewTicker(5 * time.Second)
						for {
							select {
							case <-ctx.Done():
								return
							case <-ticker.C:
								if err = json.Unmarshal([]byte(msg.Payload), &pos); err != nil {
									log.Errorf("error while unmarchaling: %v", err)
								}
								pnl := p.getPNL(pos)
								log.Infof("position [%v]%v pnl = %v", pos.Id, pos.Name, pnl)
							}
						}
					}()
					continue
				}
			}
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func NewPositionService(positionRepo *repository.PositionRepository, currentPrices *[]models.Price) *PositionService {

	service := &PositionService{
		positionRepo:  positionRepo,
		currentPrices: currentPrices,
	}

	go func() {
		err := service.ListenNotify()
		if err != nil {
			log.Errorf("unable to start listenning pg notify")
		}
	}()

	return service
}

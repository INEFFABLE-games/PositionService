package service

import (
	"context"
	"encoding/json"
	"github.com/INEFFABLE-games/PriceService/models"
	"github.com/INEFFABLE-games/PriceService/protocol"
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

type PriceService struct {
	priceChannel chan []models.Price
}

func (p *PriceService) Refresh(ctx context.Context,stream protocol.PriceService_SendClient){

	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:

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
				p.priceChannel <- butchOfPrices
			}()
		}
	}

}

func NewPriceService(priceChannel chan []models.Price)*PriceService{
	return &PriceService{priceChannel: priceChannel}
}

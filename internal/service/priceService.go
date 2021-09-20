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

// PriceService structure for PriceService objects
type PriceService struct {
	currentPrices *[]models.Price
}

// Refresh get fresh prices from grpc stream and write into currentPrices map
func (p *PriceService) Refresh(ctx context.Context, stream protocol.PriceService_SendClient) {

	ticker := time.NewTicker(1 * time.Millisecond)
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
				*p.currentPrices = butchOfPrices
			}()
		}
	}
}

// NewPriceService creates PriceService object
func NewPriceService(currentPrices *[]models.Price) *PriceService {
	return &PriceService{
		currentPrices: currentPrices,
	}
}

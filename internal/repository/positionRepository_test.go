package repository

import (
	"context"
	"database/sql"
	"github.com/INEFFABLE-games/PositionService/internal/config"
	"github.com/INEFFABLE-games/PriceService/models"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPositionRepository_Insert(t *testing.T) {

	cfg := config.NewConfig()

	conn, err := sql.Open("postgres", cfg.PsqlURI)
	if err != nil {
		log.Errorf("unable to connect with postgres %v", err)
	}

	r := NewPositionRepository(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err = r.Insert(ctx, models.Price{
		Id:   "3",
		Name: "Xiaomi",
		Bid:  1640,
		Ask:  0,
	}, "me")

	require.Nil(t, err)
}

func TestPositionRepository_Update(t *testing.T) {

	cfg := config.NewConfig()

	conn, err := sql.Open("postgres", cfg.PsqlURI)
	if err != nil {
		log.Errorf("unable to connect with postgres %v", err)
	}

	r := NewPositionRepository(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	err = r.Update(ctx, models.Price{
		Id:   "",
		Name: "Apple",
		Bid:  0,
		Ask:  1200,
	}, "me")

	require.Nil(t, err)
}

func TestPositionRepository_Get(t *testing.T) {
	cfg := config.NewConfig()

	conn, err := sql.Open("postgres", cfg.PsqlURI)
	if err != nil {
		log.Errorf("unable to connect with postgres %v", err)
	}

	r := NewPositionRepository(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := r.GetByOwner(ctx, "me")

	log.Info(res)

	require.Nil(t, err)
}

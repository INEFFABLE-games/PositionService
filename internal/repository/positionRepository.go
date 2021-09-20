package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/INEFFABLE-games/PriceService/models"
	"time"
)

type PositionRepository struct {
	db *sql.DB
}

func (p *PositionRepository) Insert(ctx context.Context, price models.Price, owner string) error {
	price.Id = fmt.Sprintf("%d", time.Now().Unix())
	_, err := p.db.ExecContext(ctx, "insert into positions(id,owner,name,ask,bid) values($1,$2,$3,$4,$5)", price.Id, owner, price.Name, price.Bid, 0)
	return err
}

func (p *PositionRepository) Update(ctx context.Context, price models.Price, owner string) error {
	_, err := p.db.ExecContext(ctx, "update positions set bid=$1 where owner=$2 and name=$3", price.Ask, owner, price.Name)
	return err
}

func (p *PositionRepository) GetByOwner(ctx context.Context, owner string) ([]models.Price, error) {
	rows, err := p.db.QueryContext(ctx, "select * from positions where owner = $1", owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// An album slice to hold data from returned rows.
	var positions []models.Price

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var pos models.Price
		var owner string
		err := rows.Scan(&pos.Id, &owner, &pos.Name, &pos.Bid, &pos.Ask)
		if err != nil {
			return positions, err
		}
		positions = append(positions, pos)
	}
	if err = rows.Err(); err != nil {
		return positions, err
	}

	return positions, nil
}

func NewPositionRepository(db *sql.DB) *PositionRepository {
	return &PositionRepository{db: db}
}

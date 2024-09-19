package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/igortoigildin/stupefied_bell/internal/model"
	"github.com/igortoigildin/stupefied_bell/internal/storage"
	"github.com/igortoigildin/stupefied_bell/pkg/logger"
	"go.uber.org/zap"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (rep *Repository) SaveOrder(ctx context.Context, order model.Order) (string, error) {
	var number string
	query := `INSERT INTO orders (number, title, quantity, comment, uploaded_at)
	VALUES ($1, $2, $3, $4, now() AT TIME ZONE 'MSK') ON CONFLICT DO NOTHING RETURNING number;`

	args := []any{
		order.Number,
		order.Title,
		order.Quantity,
		order.Comment,
	}

	err := rep.db.QueryRowContext(ctx, query, args...).Scan(&number)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return number, nil // order already exists
		default:
			logger.Log.Error("error while inserting order", zap.Error(err))
			return number, err
		}
	}
	return number, nil
}

func (rep *Repository) SelectAllOrders(ctx context.Context) ([]model.Order, error) {
	var orders []model.Order
	query := `SELECT * FROM orders;`
	rows, err := rep.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var order model.Order
		err = rows.Scan(&order.Number, &order.Quantity, &order.Title, &order.Comment, &order.UploadedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (rep *Repository) DeleteOrder(ctx context.Context, number string) error {
	_, err := rep.db.Exec("DELETE FROM orders WHERE number = $1;", number)
	if err != nil {
		return err
	}
	return nil
}

func (rep *Repository) UpdateOrder(ctx context.Context, order model.Order) error {
	query := `
	UPDATE orders SET title = $1, quantity = $2, comment = $3 WHERE number = $4;`
	res, err := rep.db.Exec(query, order.Title, order.Quantity, order.Comment, order.Number)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return storage.ErrOrderNotFound // in case no such order found, return custom error
	}
	return nil
}

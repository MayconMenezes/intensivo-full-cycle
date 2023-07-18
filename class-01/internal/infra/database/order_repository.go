package database

import (
	"database/sql"

	"github.com/maycon_menezes/go-intensivo-full-cycle/internal/entity"
)

type OrderRepository struct {
	DB *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		DB: db,
	}
}

func (o *OrderRepository) Save(order *entity.Order) error {
	_, err := o.DB.Exec("insert into orders(id, price, tax, final_price) values($1, $2, $3, $4)", order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderRepository) GetTotalTransactions() (int, error) {
	var total int
	err := o.DB.QueryRow("select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}

	return total, nil
}

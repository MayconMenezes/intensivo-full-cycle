package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) { // it is common to use pointer when we are working with go
	order := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.Validate()
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.Tax <= 0 {
		return errors.New("tax must be greater than zero")
	}

	return nil
}

func (o *Order) CalculateFinalPrice() error { // the * is a pointer receiver, that means that the method can modify the struct
	o.FinalPrice = o.Price + o.Tax

	err := o.Validate()
	if err != nil {
		return err
	}

	return nil
}

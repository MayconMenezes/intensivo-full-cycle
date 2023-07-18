package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetAnErrorIfIDIsBlank(t *testing.T) {
	_, err := NewOrder("", 10, 1)

	if err == nil {
		t.Error("it should return an error if id is blank")
	}
}

func TestIfItGetAnErrorIfIDIsBlank2(t *testing.T) {
	order := Order{}

	assert.Error(t, order.Validate(), "it should return an error if id is blank")
}

func TestFinalPrice(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, "123", order.ID)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}

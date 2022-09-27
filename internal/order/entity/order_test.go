package entity_test

import (
	"testing"

	"github.com/alecanutto/fullcycle-go/internal/order/entity"
	"github.com/stretchr/testify/assert"
)

func TestGivenAnEmptyId_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{}
  assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenAnEmptyPrice_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "1"}
  assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAnEmptyTax_WhenCreateANewOrder_ThenShouldReceiveAndError(t *testing.T) {
	order := entity.Order{ID: "1", Price: 10}
  assert.Error(t, order.IsValid(), "invalid tax")
}

func TestGivenAValidParams_WhenCallNewOrder_ThenShouldReceiveCreateOrderWithAllParams(t *testing.T) {
	order, err := entity.NewOrder("1", 10, 2)
  assert.NoError(t, err)
  assert.Equal(t, "1", order.ID)
  assert.Equal(t, 10.0, order.Price)
  assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAValidParams_WhenCallCalculateFinalPrice_ThenShouldCalculateFinalPriceAndSetItOnFinalPriceProperty(t *testing.T) {
	order, err := entity.NewOrder("1", 10, 2)
  assert.NoError(t, err)
	err = order.CalculateFinalPrice()
	assert.NoError(t, err)
	assert.Equal(t, 12.0, order.FinalPrice)
}



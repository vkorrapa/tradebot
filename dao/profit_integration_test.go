// +build integration

package dao

import (
	"testing"

	"github.com/jeremyhahn/tradebot/entity"
	"github.com/stretchr/testify/assert"
)

func TestProfitDAO(t *testing.T) {
	ctx := NewIntegrationTestContext()
	chartDAO := NewChartDAO(ctx)
	profitDAO := NewProfitDAO(ctx)

	chart := createIntegrationTestChart(ctx)
	trades := chart.GetTrades()

	chartDAO.Create(chart)

	profit := &entity.Profit{
		UserId:   ctx.GetUser().GetId(),
		TradeId:  trades[0].GetId(),
		Quantity: "1",
		Bought:   trades[0].GetPrice(),
		Sold:     trades[1].GetPrice(),
		Fee:      "2.75",
		Tax:      "5.50",
		Total:    "10008.25"}

	err := profitDAO.Create(profit)
	assert.Equal(t, nil, err)

	persisted, exErr := profitDAO.GetByTrade(&trades[0])
	assert.Equal(t, nil, exErr)
	assert.Equal(t, uint(1), persisted.GetId())
	assert.Equal(t, ctx.GetUser().GetId(), persisted.GetUserId())
	assert.Equal(t, profit.TradeId, persisted.GetTradeId())
	assert.Equal(t, profit.Quantity, persisted.GetQuantity())
	assert.Equal(t, profit.Bought, persisted.GetBought())
	assert.Equal(t, profit.Sold, persisted.GetSold())
	assert.Equal(t, profit.Fee, persisted.GetFee())
	assert.Equal(t, profit.Tax, persisted.GetTax())
	assert.Equal(t, profit.Total, persisted.GetTotal())

	persisted2, exErr2 := profitDAO.Find()
	assert.Equal(t, nil, exErr2)
	assert.Equal(t, 1, len(persisted2))
	assert.Equal(t, uint(1), persisted2[0].GetId())
	assert.Equal(t, ctx.GetUser().GetId(), persisted2[0].GetUserId())
	assert.Equal(t, profit.TradeId, persisted2[0].GetTradeId())
	assert.Equal(t, profit.Quantity, persisted2[0].GetQuantity())
	assert.Equal(t, profit.Bought, persisted2[0].GetBought())
	assert.Equal(t, profit.Sold, persisted2[0].GetSold())
	assert.Equal(t, profit.Fee, persisted2[0].GetFee())
	assert.Equal(t, profit.Tax, persisted2[0].GetTax())
	assert.Equal(t, profit.Total, persisted2[0].GetTotal())

	CleanupIntegrationTest()
}

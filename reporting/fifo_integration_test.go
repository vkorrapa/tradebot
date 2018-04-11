package reporting

import (
	"testing"
	"time"

	"github.com/jeremyhahn/tradebot/test"
	"github.com/stretchr/testify/assert"
)

func TestFIFO(t *testing.T) {
	ctx := test.NewIntegrationTestContext()

	transactions, err := createTransactionService(ctx).GetHistory()
	assert.Nil(t, err)

	fifo := NewFifoReport(ctx, transactions)
	location := time.Now().Location()
	start := time.Date(2017, 01, 01, 0, 0, 0, 0, location)
	end := time.Date(2017, 12, 31, 0, 0, 0, 0, location)
	form8948 := fifo.Run(start, end)
	form8948.WriteCSV("/tmp/8949.csv")

	//test.CleanupIntegrationTest()
}

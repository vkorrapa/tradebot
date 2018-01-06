package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/jeremyhahn/tradebot/common"
	logging "github.com/op/go-logging"
)

type Ripple struct {
	logger *logging.Logger
	client http.Client
	common.Wallet
}

type RippleWallet struct {
	Balance  string `json:"value"`
	Currency string `json:"currency"`
}

type RippleBalance struct {
	Result      string         `json:"result"`
	LedgerIndex float64        `json:"ledger_index"`
	Limit       int64          `json:"limit"`
	Balances    []RippleWallet `json:"balances"`
}

func NewRipple(ctx *common.Context) *Ripple {
	client := http.Client{
		Timeout: time.Second * 2}
	return &Ripple{
		logger: ctx.Logger,
		client: client}
}

func (r *Ripple) GetBalance(address string) *common.CryptoWallet {

	r.logger.Debugf("[Ripple.GetBalance] Address: %s", address)

	var balance float64
	url := fmt.Sprintf("https://data.ripple.com/v2/accounts/%s/balances", address)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		r.logger.Errorf("[Ripple.GetBalance] %s", err.Error())
	}

	req.Header.Set("User-Agent", fmt.Sprintf("%s/v%s", common.APPNAME, common.APPVERSION))

	res, getErr := r.client.Do(req)
	if getErr != nil {
		r.logger.Errorf("[Ripple.GetBalance] %s", getErr.Error())
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		r.logger.Errorf("[Ripple.GetBalance] %s", readErr.Error())
	}

	rb := RippleBalance{}
	jsonErr := json.Unmarshal(body, &rb)
	if jsonErr != nil {
		r.logger.Errorf("[Ripple.GetBalance] %s", jsonErr.Error())
	}
	if len(rb.Balances) <= 0 {
		balance = 0.0
	} else {
		f, _ := strconv.ParseFloat(rb.Balances[0].Balance, 64)
		balance = f
	}

	//marketcap := NewMarketCapService(ctx, )

	return &common.CryptoWallet{
		Address:  address,
		Balance:  balance,
		Currency: "XRP"}
	// NetWorth: }
}
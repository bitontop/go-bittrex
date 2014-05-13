package bittrex

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const (
	API_BASE                   = "https://bittrex.com/api/" // Bittrex API endpoint
	API_VERSION                = "v1.1"                     // Bittrex API version
	DEFAULT_HTTPCLIENT_TIMEOUT = 30                         // HTTP client timeout
)

// bittrex represent a bittrex client
type bittrex struct {
	client *client
}

// New return a instanciate bittrex struct
func New(apiKey, apiSecret string) *bittrex {
	client := NewClient(apiKey, apiSecret)
	return &bittrex{client}
}

// handleErr gets JSON response from Bittrex API en deal with error
func handleErr(r jsonResponse) error {
	if !r.Success {
		return errors.New(r.Message)
	}
	return nil
}

// GetMarkets is used to get the open and available trading markets at Bittrex along with other meta data.
func (b *bittrex) GetMarkets() (markets []Market, err error) {
	r, err := b.client.do("GET", "public/getmarkets", "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &markets)
	return
}

// GetTicker is used to get the current ticker values for a market.
func (b *bittrex) GetTicker(market string) (ticker Ticker, err error) {
	r, err := b.client.do("GET", "public/getticker?market="+strings.ToUpper(market), "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &ticker)
	return
}

// GetMarketSummaries is used to get the last 24 hour summary of all active exchanges
func (b *bittrex) GetMarketSummaries() (marketSummaries []MarketSummary, err error) {
	r, err := b.client.do("GET", "public/getmarketsummaries", "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &marketSummaries)
	return
}

// GetOrderBook is used to get retrieve the orderbook for a given market
// market: a string literal for the market (ex: BTC-LTC)
// cat: buy, sell or both to identify the type of orderbook to return.
// depth: how deep of an order book to retrieve. Max is 100
func (b *bittrex) GetOrderBook(market, cat string, depth int) (orderBook OrderBook, err error) {
	if cat != "buy" && cat != "sell" && cat != "both" {
		cat = "both"
	}
	if depth > 100 {
		depth = 100
	}
	if depth < 1 {
		depth = 1
	}
	r, err := b.client.do("GET", fmt.Sprintf("public/getorderbook?market=%s&type=%s&depth=%d", strings.ToUpper(market), cat, depth), "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &orderBook)
	return
}

// GetMarketHistory is used to retrieve the latest trades that have occured for a specific market.
// mark a string literal for the market (ex: BTC-LTC)
// count a number between 1-100 for the number of entries to return
func (b *bittrex) GetMarketHistory(market string, count int) (trades []Trade, err error) {
	if count > 100 {
		count = 100
	}
	if count < 1 {
		count = 1
	}
	r, err := b.client.do("GET", fmt.Sprintf("public/getmarkethistory?market=%s&count=%d", strings.ToUpper(market), count), "", false)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &trades)
	return
}

// Market

// BuyLimit is used to place a limited buy order in a specific market.
func (b *bittrex) BuyLimit(market string, quantity, rate float64) (uuid string, err error) {
	r, err := b.client.do("GET", "market/buylimit?market="+market+"&quantity="+strconv.FormatFloat(quantity, 'f', 8, 64)+"&rate="+strconv.FormatFloat(rate, 'f', 8, 64), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var u Uuid
	err = json.Unmarshal(response.Result, &u)
	uuid = u.Id
	return
}

// BuyMarket is used to place a market buy order in a spacific market.
func (b *bittrex) BuyMarket(market string, quantity float64) (uuid string, err error) {
	r, err := b.client.do("GET", "market/buymarket?market="+market+"&quantity="+strconv.FormatFloat(quantity, 'f', 8, 64), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var u Uuid
	err = json.Unmarshal(response.Result, &u)
	uuid = u.Id
	return
}

// SellLimit is used to place a limited sell order in a specific market.
func (b *bittrex) SellLimit(market string, quantity, rate float64) (uuid string, err error) {
	r, err := b.client.do("GET", "market/selllimit?market="+market+"&quantity="+strconv.FormatFloat(quantity, 'f', 8, 64)+"&rate="+strconv.FormatFloat(rate, 'f', 8, 64), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var u Uuid
	err = json.Unmarshal(response.Result, &u)
	uuid = u.Id
	return
}

// SellMarket is used to place a market sell order in a specific market.
func (b *bittrex) SellMarket(market string, quantity float64) (uuid string, err error) {
	r, err := b.client.do("GET", "market/selllimit?market="+market+"&quantity="+strconv.FormatFloat(quantity, 'f', 8, 64), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var u Uuid
	err = json.Unmarshal(response.Result, &u)
	uuid = u.Id
	return
}

// CancelOrder is used to cancel a buy or sell order.
func (b *bittrex) CancelOrder(orderID string) (err error) {
	r, err := b.client.do("GET", "market/cancel?uuid="+orderID, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	err = handleErr(response)
	return
}

// GetOpenOrders returns orders that you currently have opened.
// If market is set to "all", GetOpenOrders return all orders
// If market is set to a specific order, GetOpenOrders return orders for this market
func (b *bittrex) GetOpenOrders(market string) (openOrders []Order, err error) {
	ressource := "market/getopenorders"
	if market != "all" {
		ressource += "?market=" + strings.ToUpper(market)
	}
	r, err := b.client.do("GET", ressource, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &openOrders)
	return
}

// Account

// GetBalances is used to retrieve all balances from your account
func (b *bittrex) GetBalances() (balances map[string]Balance, err error) {
	r, err := b.client.do("GET", "account/getbalances", "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &balances)
	return
}

// Getbalance is used to retrieve the balance from your account for a specific currency.
// currency: a string literal for the currency (ex: LTC)
func (b *bittrex) GetBalance(currency string) (balance Balance, err error) {
	r, err := b.client.do("GET", "account/getbalance?currency="+strings.ToUpper(currency), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &balance)
	return
}

// GetDepositAddress is sed to generate or retrieve an address for a specific currency.
// currency a string literal for the currency (ie. BTC)
func (b *bittrex) GetDepositAddress(currency string) (address Address, err error) {
	r, err := b.client.do("GET", "account/getdepositaddress?currency="+strings.ToUpper(currency), "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &address)
	return
}

// Withdraw is used to withdraw funds from your account.
// address string the address where to send the funds.
// currency string literal for the currency (ie. BTC)
// quantity float the quantity of coins to withdraw
func (b *bittrex) Withdraw(address, currency string, quantity float64) (withdrawUuid string, err error) {
	r, err := b.client.do("GET", "account/withdraw?currency="+strings.ToUpper(currency)+"&quantity="+strconv.FormatFloat(quantity, 'f', 8, 64)+"&address="+address, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	var u Uuid
	err = json.Unmarshal(response.Result, &u)
	withdrawUuid = u.Id
	return
}

// GetOrderHistory used to retrieve your order history.
// market string literal for the market (ie. BTC-LTC). If set to "all", will return for all market
// count int : 	the number of records to return. Is set to 0, will return max history
func (b *bittrex) GetOrderHistory(market string, count int) (orders []Order, err error) {
	ressource := "account/getorderhistory"
	if count != 0 || market != "all" {
		ressource += "?"
	}
	if count != 0 {
		ressource += fmt.Sprintf("count=%d&", count)
	}
	if market != "all" {
		ressource += "market=" + market
	}
	r, err := b.client.do("GET", ressource, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &orders)
	return
}

// GetWithdrawalHistory is used to retrieve your withdrawal history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
// count int the number of records to return. Is set to 0 will return the max set.
func (b *bittrex) GetWithdrawalHistory(currency string, count int) (withdrawals []Withdrawal, err error) {
	ressource := "account/getwithdrawalhistory"
	if count != 0 || currency != "all" {
		ressource += "?"
	}
	if count != 0 {
		ressource += fmt.Sprintf("count=%d&", count)
	}
	if currency != "all" {
		ressource += "currency=" + currency
	}
	r, err := b.client.do("GET", ressource, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &withdrawals)
	return
}

// GetDepositHistory is used to retrieve your deposit history
// currency string a string literal for the currency (ie. BTC). If set to "all", will return for all currencies
// count int the number of records to return. Is set to 0 will return the max set.
func (b *bittrex) GetDepositHistory(currency string, count int) (deposits []Deposit, err error) {
	ressource := "account/getdeposithistory"
	if count != 0 || currency != "all" {
		ressource += "?"
	}
	if count != 0 {
		ressource += fmt.Sprintf("count=%d&", count)
	}
	if currency != "all" {
		ressource += "currency=" + currency
	}
	r, err := b.client.do("GET", ressource, "", true)
	if err != nil {
		return
	}
	var response jsonResponse
	if err = json.Unmarshal(r, &response); err != nil {
		return
	}
	if err = handleErr(response); err != nil {
		return
	}
	err = json.Unmarshal(response.Result, &deposits)
	return
}

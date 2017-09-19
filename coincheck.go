package coincheck

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CoinCheck struct {
	AccessKey   string
	SecretKey   string
	Account     Account
	BankAccount BankAccount
	Borrow      Borrow
	Deposit     Deposit
	Leverage    Leverage
	Order       Order
	OrderBook   OrderBook
	Send        Send
	Ticker      Ticker
	Trade       Trade
	Transfer    Transfer
	Withdraw    Withdraw
}

func (c CoinCheck) NewClient(accessKey string, secretKey string) CoinCheck {
	c.AccessKey = accessKey
	c.SecretKey = secretKey
	c.Account = Account{&c}
	c.BankAccount = BankAccount{&c}
	c.Borrow = Borrow{&c}
	c.Deposit = Deposit{&c}
	c.Leverage = Leverage{&c}
	c.Order = Order{&c}
	c.OrderBook = OrderBook{&c}
	c.Send = Send{&c}
	c.Ticker = Ticker{&c}
	c.Trade = Trade{&c}
	c.Transfer = Transfer{&c}
	c.Withdraw = Withdraw{&c}
	return c
}

func (c CoinCheck) Request(method string, path string, param string) string {
	if param != "" && method == "GET" {
		path = path + "?" + param
		param = ""
	}
	url := "https://coincheck.jp/" + path
	nonce := strconv.FormatInt(CreateNonce(), 10)
	message := nonce + url + param
	req := &http.Request{}
	if method == "POST" {
		payload := strings.NewReader(param)
		req, _ = http.NewRequest(method, url, payload)
	} else {
		req, _ = http.NewRequest(method, url, nil)
	}
	signature := ComputeHmac256(message, c.SecretKey)
	req.Header.Add("access-key", c.AccessKey)
	req.Header.Add("access-nonce", nonce)
	req.Header.Add("access-signature", signature)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return string(body)
}

//create nonce by milliseconds
func CreateNonce() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond) / int64(time.Nanosecond))
}

//create signature
func ComputeHmac256(message string, secret string) string {
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return hex.EncodeToString(h.Sum(nil))
}

package coincheck_test

import (
	"github.com/Akagi201/coincheckgo"
)

func ExampleClient() {
	client := new(coincheck.CoinCheck).NewClient("CTBjkbYihKmT-IHR", "DmkP3ChVjOjqieaX-jFnU4XxpSlSgT3M")
	/** Public API */
	client.Ticker.All()
	client.Trade.All()
	client.OrderBook.All()

	/** Private API */
	// 新規注文
	// "buy" 指値注文 現物取引 買い
	// "sell" 指値注文 現物取引 売り
	// "market_buy" 成行注文 現物取引 買い
	// "market_sell" 成行注文 現物取引 売り
	// "leverage_buy" 指値注文 レバレッジ取引新規 買い
	// "leverage_sell" 指値注文 レバレッジ取引新規 売り
	// "close_long" 指値注文 レバレッジ取引決済 売り
	// "close_short" 指値注文 レバレッジ取引決済 買い
	client.Order.Create(`{"rate":"28500","amount":"0.00508771", "order_type":"buy", "pair":"btc_jpy"}`)
	// 未決済の注文一覧
	client.Order.Opens()
	// 注文のキャンセル
	client.Order.Cancel("12345")
	// 取引履歴
	client.Order.Transactions()
	// ポジション一覧
	client.Leverage.Positions()
	// 残高
	client.Account.Balance()
	// レバレッジアカウントの残高
	client.Account.LeverageBalance()
	// アカウント情報
	client.Account.Info()
	// ビットコインの送金
	client.Send.Create(`{"address":"1Gp9MCp7FWqNgaUWdiUiRPjGqNVdqug2hY","amount":"0.0002"`)
	// ビットコインの送金履歴
	client.Send.All("currency=BTC")
	// ビットコインの受け取り履歴
	client.Deposit.All("currency=BTC")
	// ビットコインの高速入金
	client.Deposit.Fast("12345")
	// 銀行口座一覧
	client.BankAccount.All()
	// 銀行口座の登録
	client.BankAccount.Create(`{"bank_name":"MUFG","branch_name":"tokyo", "bank_account_type":"toza", "number":"1234567", "name":"Danny"}`)
	// 銀行口座の削除
	client.BankAccount.Delete("25621")
	// 出金履歴
	client.Withdraw.All()
	// 出金申請の作成
	client.Withdraw.Create(`{"bank_account_id":"2222","amount":"50000", "currency":"JPY", "is_fast":"false"}`)
	// 出金申請のキャンセル
	client.Withdraw.Cancel("12345")
	// 借入申請
	client.Borrow.Create(`{"amount":"100","currency":"JPY"}`)
	// 借入中一覧
	client.Borrow.Matches()
	// 返済
	client.Borrow.Repay("1135")
	// レバレッジアカウントへの振替
	client.Transfer.ToLeverage(`{"amount":"100","currency":"JPY"}`)
	// レバレッジアカウントからの振替
	client.Transfer.FromLeverage(`{"amount":"100","currency":"JPY"}`)
}

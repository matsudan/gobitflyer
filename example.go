package main

import (
	"context"
	"fmt"
	"github.com/matsudan/gobitflyer/bitflyer"
)

func main() {
	cfg := bitflyer.LoadConfig()
	client := bitflyer.NewClient(*cfg)

	ctx := context.Background()

	//outputMarkets, err := client.GetMarketList(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("market list: %#v\n", outputMarkets)

	//outputBoard, err := client.GetBoard(ctx, "BTC_JPY")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("output board: %#v\n", outputBoard)

	outputTicker, err := client.GetTicker(ctx, "BTC_JPY")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("ticker: %#v\n", outputTicker)

	//pq := bitflyer.PaginationQuery{
	//	Count: "10",
	//}
	//outputExecutions, err := client.GetExecutionList(ctx, "BTC_JPY", pq)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("executions: %#v\n", outputExecutions)

	//outputBoardState, err := client.GetBoardState(ctx, "BTC_JPY")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("board state: %#v\n", outputBoardState)

	//outputChats, err := client.GetChatList(ctx, "2020-10-05")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("chats: %#v\n", outputChats.Chats[0])

	//outputBalance, err := client.GetBalanceList(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("balance list: %#v\n",outputBalance.Balance[0])

	//outputCollateral, err := client.GetCollateral(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("collateral: %#v\n",outputCollateral)

	//outputCollateralAccounts, err := client.GetCollateralAccountList(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("collateral accounts: %#v\n",outputCollateralAccounts.CollateralAccounts)

	outputAddresses, err := client.GetAddressList(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("addresses: %#v\n",outputAddresses.Addresses)
}

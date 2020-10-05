package main

import (
	"context"
	"fmt"

	"github.com/matsudan/gobitflyer/lightning/public"
)

func main() {
	options := public.Options{
		Version: "v1",
	}
	client, _ := public.New(options)

	ctx := context.Background()

	outputMarkets, err := client.GetMarketList(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("market list: %#v\n", outputMarkets)

	//outputBoard, err := client.GetBoard(ctx, "BTC_JPY")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("output board: %#v\n", outputBoard)

	//outputTicker, err := client.GetTicker(ctx, "BTC_JPY")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("ticker: %#v\n", outputTicker)

	//pq := public.PaginationQuery{
	//	Count: "10",
	//}
	//
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

	//outputChats, err := client.GetChatList(ctx, "2020-10-01")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("chats: %#v\n", outputChats[0])
}

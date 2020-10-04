package main

import (
	"context"
	"fmt"

	"github.com/matsudan/gobitflyer/lightning"
)

func main() {
	options := lightning.Options{
		Version: "v1",
	}
	client, _ := lightning.New(options)

	ctx := context.Background()

	outputMarkets, err := client.GetMarketList(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("market list: %#v\n", outputMarkets[0])

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
	//
	//pq := lightning.PaginationQuery{
	//	Count: "10",
	//}

	//outputExecutions, err := client.GetExecutionList(ctx, "BTC_JPY", pq)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Printf("executions: %#v\n", outputExecutions[0])

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

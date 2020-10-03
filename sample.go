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
	outputHealth, err := client.GetHealth(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("output health: %#v\n", *outputHealth)

	outputMarkets, err := client.GetMarketList(ctx)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("output market list: %#v\n", outputMarkets[0])
}


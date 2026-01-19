package main

import (
	"context"
	"fmt"
	"time"
)

func AggregateStockPrice(symbol string, sources []StockSource, timeout time.Duration) (StockPrice, error) {
	// TODO: Implement this function

	panic("not implemented")
}

func AggregateStockPriceWithContext(ctx context.Context, symbol string, sources []StockSource) (StockPrice, error) {
	// TODO: Implement this function

	panic("not implemented")
}

func main() {
	sources := GetSources()

	fmt.Println("Fetching AAPL stock price from multiple sources...")

	start := time.Now()
	price, err := AggregateStockPrice("AAPL", sources, 5*time.Second)
	elapsed := time.Since(start)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Got price $%.2f from %s in %v\n", price.Price, price.Source, elapsed)
}

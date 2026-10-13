package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	adanos "github.com/adanos-software/adanos-go-sdk"
)

func main() {
	c, err := adanos.NewClient(adanos.Config{APIKey: os.Getenv("ADANOS_API_KEY")})
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	health, err := c.Health(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(health.Status, health.Version)
	if os.Getenv("ADANOS_API_KEY") == "" {
		return
	}
	stocks, err := c.Reddit().Trending(ctx, &adanos.TrendingOptions{Limit: new(3)})
	if err != nil {
		log.Fatal(err)
	}
	for _, stock := range stocks {
		fmt.Println(stock.Ticker, stock.BuzzScore)
	}
}

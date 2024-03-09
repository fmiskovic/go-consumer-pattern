package main

import (
	"fmt"

	"github.com/fmiskovic/go-consumer-pattern/internal/bar"
	"github.com/fmiskovic/go-consumer-pattern/internal/beer"
	"github.com/fmiskovic/go-consumer-pattern/internal/whisky"
	"github.com/fmiskovic/go-consumer-pattern/internal/wine"
)

func main() {
	bar := bar.Bar{}
	// beer consumer
	beer := beer.New(bar)
	fmt.Println(beer.Order())
	// wine consumer
	wine := wine.New(bar)
	fmt.Println(wine.Order())
	// whisky consumer
	whisky := whisky.New(bar)
	fmt.Println(whisky.Order())
}

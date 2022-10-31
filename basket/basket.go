package main

import (
	"fmt"
)

type Item struct {
	title, description string
	quantity, pricePer float32
}

func main() {
	item1 := Item{
		title:       "Lego Set",
		description: "4000 pieces",
		quantity:    1,
		pricePer:    600.00,
	}

	item2 := Item{
		title:       "Plush",
		description: "plush toy",
		quantity:    3,
		pricePer:    5.99,
	}

	items := make([]Item, 0)
	items = append(items, item1)
	items = append(items, item2)

	var sum float32 = 0
	for _, item := range items {
		sum += item.quantity * item.pricePer
		fmt.Println("Found ", item.quantity, item.title)
	}

	fmt.Println("Cart total", sum)
}

package main

import (
	"105/food"
	"fmt"
)

var m = map[string]food.MenuItem{
	"Main": {
		"spaghetti", 14.99,
	},
	"Dessert": {
		"spumoni", 22.49,
	},
}

func main() {
	food.PizzanosMenu()
	donutMenu()
	italianMenu()
}

func italianMenu() {
	fmt.Println("Francesco's")
	for a, b := range m {
		fmt.Printf("%v\n---\n%v $%v\n", a, b.Desc, b.Price)
	}
}

func donutMenu() {
	menu := make(map[string]float64)
	menu["bear claw"] = 2.50
	menu["glazed"] = 1.25
	fmt.Println("Donut Holz")
	for dish, price := range menu {
		fmt.Println(dish, price)
	}
}

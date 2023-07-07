package main

import (
	"fmt"
	"github.com/pgiles/105/food/enum"

	"github.com/pgiles/105/food"
	"github.com/pgiles/105/openapi"
)

var m = map[string]food.MenuItem{
	"Main": {
		"spaghetti", 14.99, enum.Undefined,
	},
	"Dessert": {
		"spumoni", 22.49, enum.Undefined,
	},
}

func main() {
	food.PizzanosMenu()
	donutMenu()
	italianMenu()

	b, err := openapi.Generate("a.cue", &openapi.Info{
		Desc: "let's find out",
	})
	if err != nil {
		panic(err)
	}
	fmt.Print(string(b))
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

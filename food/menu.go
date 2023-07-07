package food

import (
	"fmt"
	"github.com/pgiles/105/food/enum"
)

// Menu has a description and list of dishes
type Menu struct {
	Desc  string
	Items []MenuItem
}

// MenuItem are the dishes of a menu
type MenuItem struct {
	Desc  string
	Price float64
	Promo enum.Promotion
}

// PizzanosMenu static default menu
func PizzanosMenu() Menu {
	return createMenu()
	//menu.Print()
}

func createMenu() Menu {
	menuItem1 := MenuItem{
		Desc:  "Pizza",
		Price: 23.99,
	}
	menuItem2 := MenuItem{
		Desc:  "Pasta",
		Price: 16.99,
	}
	menuItem3 := MenuItem{
		Desc:  "Baked Ziti",
		Price: 18.99,
		Promo: enum.Daily,
	}
	menuItem4 := MenuItem{
		Desc:  "Summer Bowtie",
		Price: 13.99,
		Promo: enum.Seasonal,
	}
	menu := Menu{
		Desc:  "Pizzanos Daily Menu",
		Items: []MenuItem{menuItem1, menuItem2, menuItem3, menuItem4},
	}
	return menu
}

func (m *Menu) Print() {
	fmt.Println(m.Desc)
	for _, itm := range m.Items {
		promoIndicator := ""
		if itm.Promo == enum.Daily {
			promoIndicator = "*"
		} else if itm.Promo == enum.Seasonal {
			promoIndicator = ":-)"
		}
		fmt.Println(itm.Desc, ":", itm.Price, promoIndicator)

	}
}

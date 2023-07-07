package food

import "fmt"

// Menu has a description and list of dishes
type Menu struct {
	Desc  string
	Items []MenuItem
}

// MenuItem are the dishes of a menu
type MenuItem struct {
	Desc  string
	Price float64
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
	menu := Menu{
		Desc:  "Pizzanos",
		Items: []MenuItem{menuItem1, menuItem2},
	}
	return menu
}

func (m *Menu) Print() {
	fmt.Println(m.Desc)
	for _, itm := range m.Items {
		fmt.Println(itm.Desc, ":", itm.Price)

	}
}

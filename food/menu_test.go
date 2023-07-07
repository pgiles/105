package food

import (
	"testing"
)

func TestMenu(t *testing.T) {
	menu := PizzanosMenu()
	menu.Print()
}

package main

import "fmt"

type pizza interface {
	getPrice() int
}

type base struct {
}

func (b *base) getPrice() int {
	return 15
}

type tomatoTopping struct {
	pizza
}

func (t *tomatoTopping) getPrice() int {
	pizzaPrice := t.pizza.getPrice()
	return pizzaPrice + 10
}

type cheeseTopping struct {
	pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 20
}

func main() {

	pizza := &base{}

	//Add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	//Add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("price is %d\n", pizzaWithCheeseAndTomato.getPrice())

	// Output:
	// price is 45
}

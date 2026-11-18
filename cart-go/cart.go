package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Qtr   int64
}

type Cart struct {
	item []Product
}

func (c *Cart) total() float64 {
	var sum float64 = 0
	for _, value := range c.item {
		sum += value.Price * float64(value.Qtr)
	}
	return sum
}

func (p *Product) total() float64 {
	return p.Price * float64(p.Qtr)
}

func (c *Cart) printReceipt() {
	fmt.Println("====Receipt====")
	for _, value := range c.item {
		fmt.Printf("%-6s %.2f x %d = %.2f\n",
			value.Name, value.Price, value.Qtr, value.total())
	}
	fmt.Println("------------")
	fmt.Printf("Total %.2f\n", c.total())
}

func main() {
	cart := Cart{
		item: []Product{
			{Name: "Big", Price: 200, Qtr: 4},
			{Name: "Big", Price: 400, Qtr: 5},
			{Name: "Small", Price: 500, Qtr: 50},
		},
	}
	cart.printReceipt()
}

package main

import "fmt"

type Info struct {
	Name      string
	Catergory string
	Price     float64
	Qty       int
}

type Store struct {
	Item []Info
}

func (i *Info) value() float64 {
	return i.Price * float64(i.Qty)
}

func (s *Store) Total() float64 {
	var total float64
	for i := range s.Item {
		value := s.Item[i]
		total += value.value()
	}
	return total
}

func (s *Store) Restock(name string, n int) {
	for i, v := range s.Item {
		news := &s.Item[i]
		if name == v.Name {
			news.Qty += n
		}
	}
}

func (s *Store) CountByCategory() map[string]int {
	m := map[string]int{}
	for _, v := range s.Item {
		m[v.Catergory]++
	}
	return m
}

func (s *Store) Report() {
	for _, v := range s.Item {
		fmt.Printf("%s %.2f x %d = %.2f\n", v.Name, v.Price, v.Qty, v.value())
	}
	fmt.Printf("Total %.2f\n", s.Total())
	fmt.Println(s.CountByCategory())
}

func main() {
	list := &Store{
		[]Info{
			{Name: "Pen", Catergory: "stationery", Price: 10, Qty: 20},
			{Name: "book", Catergory: "stationery", Price: 150, Qty: 5},
			{Name: "Cola", Catergory: "drink", Price: 15, Qty: 30},
		},
	}
	list.Restock("book", 5)
	list.Report()
}

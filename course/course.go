package course

import "fmt"

type course struct {
	name    string
	price   float64
	isFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func New(name string, price float64, isFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:    name,
		price:   price,
		isFree:  isFree,
	}
}

func (c *course) PrintClasses() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}

	fmt.Println(text[:len(text)-2])
}

func (c *course) ChangePrice(price float64) {
	c.price = price
}

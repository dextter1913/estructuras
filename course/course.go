package course

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func NewCourse(name string, price float64, IsFree bool) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:   name,
		Price:  price,
		IsFree: IsFree,
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
	c.Price = price
}

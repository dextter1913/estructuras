package course

import "fmt"

type course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func New(name string, price float64, isFree bool, userIds []uint, clases map[uint]string) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		Name:    name,
		Price:   price,
		IsFree:  isFree,
		UserIDs: userIds,
		Clases:  clases,
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

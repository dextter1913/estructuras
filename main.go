package main

import "fmt"

func main() {
	// Go := Course{
	// 	"Go desde Cero",
	// 	12.34,
	// 	false,
	// 	[]uint{12, 56, 89},
	// 	map[uint]string{
	// 		1: "Introduccion",
	// 		2: "Estructuras",
	// 		3: "Maps",
	// 	},
	// }

	Go := &Course{
		"Go desde Cero",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introduccion",
			2: "Estructuras",
			3: "Maps",
		},
	}

	// Go.PrintClasses()
	Go.ChangePrice(67.12)
	// (&Go).ChangePrice(67.12)
	fmt.Println(Go)
	// Go.PrintClasses()
}

package main

import (
	"fmt"
	"github.com/dextter1913/estructuras/course"
)

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

	Go := &course.Course{
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

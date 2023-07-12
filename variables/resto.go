package variables

import (
	"fmt"
	"strconv"
	"time"
)

func RestoVariables() {
	var Nombre string
	var Estado bool
	var Sueldo float32
	var Fecha time.Time

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.96
	Fecha = time.Now()

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
}

func ConvertirAString(num int) (bool, string) {
	return true, strconv.Itoa(num)
}

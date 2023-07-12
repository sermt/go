package main

import (
	"fmt"

	"github.com/sermt/go/variables"
)

func main() {
	// variables.MuestroEnteros()
	//variables.RestoVariables()
	valor, cadena := variables.ConvertirAString(88)
	fmt.Println(valor, cadena)
}

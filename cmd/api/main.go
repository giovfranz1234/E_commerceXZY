// package main
package main

import (
	"fmt"

	"github.com/giovfranz1234/E_commerceXZY/internal/server"
)

func main() {
	//Aqui se puede crear el servidor y configurarlo con las opciones deseadas
	srv := server.New(":8080")

	fmt.Println("Hello, World!")

}

package main

import (
	"log"

	"github.com/devtimx/backTwit/bd"
	"github.com/devtimx/backTwit/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la DB")
		return
	}
	handlers.Manejadores()
}

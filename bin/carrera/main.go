package main

import (
	"log"

	"github.com/ktnyt/carrera"
)

func main() {
	if err := carrera.Run(); err != nil {
		log.Fatal(err)
	}
}

package main

import (
	"log"

	"github.com/almiskov/ambilight/internal/ambilight"
)

func main() {
	cfg := ambilight.Config{
		X:       17,
		Y:       9,
		Depth:   50,
		Display: 0,
		COM:     "COM4",
	}

	err := ambilight.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}

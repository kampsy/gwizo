package main

import (
	"math/rand"
	"time"
)

func randomColour() string {
	colours := []string{
		"red", "green", "brown", "orange", "purple", "pink",
	}
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(len(colours) - 1)
	return colours[num]
}

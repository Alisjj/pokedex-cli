package main

import (
	"math/rand"
	"time"
)

func NewPokedex() Pokedex {
	a := Pokedex{
		pokemon: make(map[string]Pokemon),
	}

	return a
}

func catch(baseExp int) bool {

	// rand.Seed(time.Now().UnixNano())
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	maxExp := 200

	catchProbabilty := 100 - (baseExp * 100 / maxExp)
	catch := r.Intn(100)

	return catch < catchProbabilty

}

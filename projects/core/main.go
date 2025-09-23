package main

import "panelium/core/internal/server/router"

func main() {
	r := router.Init()

	err := r.Run()
	if err != nil {
		panic(err) //TODO: error handling
	}
}

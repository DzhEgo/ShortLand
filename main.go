package main

import "ShortLand/internal/portal/app"

func main() {
	if err := app.StartApp(); err != nil {
		panic(err)
	}
}

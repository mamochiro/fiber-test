package main

import (
	"fiber-test/internal/app"
	"log"
)

func main() {
	a := app.SetupApp()
	log.Fatal(a.Listen(":3000"))
}

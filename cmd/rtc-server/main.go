package main

import "github.com/QBasy/bailanys-server-go/internal/app"

func main() {
	application := app.NewApp()
	application.Run()
}

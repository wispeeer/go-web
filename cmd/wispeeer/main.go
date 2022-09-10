package main

import "github.com/wispeeer/go-web/internal/app"

func init() {
	//info.Version.Print()
}

func main() {
	if err := app.App(); err != nil {
		panic(err)
	}
}

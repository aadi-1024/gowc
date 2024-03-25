package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

type App struct {
	Fd *os.File
	//byte count
	C bool
}

func main() {
	//Appwide config
	app := App{}

	//filepath check
	filePath := os.Args[len(os.Args)-1]
	fd, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Invalid filePath")
		os.Exit(FAILURE)
	}
	app.Fd = fd

	//byte count
	flag.BoolVar(&app.C, "c", true, "gowc -c pathToFile")

	flag.Parse()

	counts, err := app.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(FAILURE)
	}
	fmt.Println(counts)
	os.Exit(SUCCESS)
}

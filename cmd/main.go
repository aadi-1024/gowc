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
	//Line count
	L bool
	//Word count
	W bool
	//Character count
	M bool
	//max buffer size
	Max_buf int
}

func main() {
	//Appwide config
	app := App{}

	//byte count
	flag.BoolVar(&app.C, "c", false, "gowc -c pathToFile")
	flag.BoolVar(&app.L, "l", false, "gowc -l pathToFile")
	flag.BoolVar(&app.W, "w", false, "gowc -w pathToFile")
	flag.BoolVar(&app.M, "m", false, "gowc -m pathToFile")
	flag.IntVar(&app.Max_buf, "b", 1048576, "gowc -w -b=1024 pathToFile") //1mb default
	flag.Parse()

	//in case no flags provided
	if flag.NFlag() == 0 {
		app.C = true
		app.W = true
		app.L = true
	}
	//whether file provided
	if flag.NArg() == 0 {
		app.Fd = os.Stdin
	} else {
		filePath := os.Args[len(os.Args)-1]
		fd, err := os.Open(filePath)
		if err != nil {
			fmt.Println("Invalid filePath")
			os.Exit(FAILURE)
		}
		app.Fd = fd
	}

	counts, err := app.Generate()
	if err != nil {
		fmt.Println(err)
		os.Exit(FAILURE)
	}
	fmt.Println(counts)
	os.Exit(SUCCESS)
}

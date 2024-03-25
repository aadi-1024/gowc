package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

func main() {
	//filepath check
	filePath := os.Args[len(os.Args)-1]
	fd, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Invalid filePath")
		os.Exit(FAILURE)
	}

	//byte count
	c := flag.Bool("c", false, "gowc -c filename")

	flag.Parse()
	var bytes int64

	if *c {
		fileinfo, err := fd.Stat()
		if err != nil {
			fmt.Println(err)
			os.Exit(FAILURE)
		}
		bytes = fileinfo.Size()
	}
	fmt.Printf("%v %v\n", bytes, filepath.Base(filePath))
	os.Exit(SUCCESS)
}

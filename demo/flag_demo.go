package main

import (
	"flag"
	"fmt"
	"os"
)

// go run main.go -data jalins
//func main()  {
//	data := flag.String("data", "jalins", "this is a string")
//	flag.Parse()
//
//	fmt.Println(*data)
//
//}

// go run main.go addData -data jalins

func main()  {
	newFlagSet := flag.NewFlagSet("addData", flag.ExitOnError)

	data := newFlagSet.String("data", "jalins", "addData string")

	fmt.Println(newFlagSet.Parse(os.Args[2:]))

	if newFlagSet.Parsed() {
		if *data == "" {
			newFlagSet.Usage()
		}
	}

}
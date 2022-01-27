package main

import (
	"flag"
	"fmt"
)

func main() {
	stringPtr := flag.String("word", "random text", "test for string") // flag.<dataType>(<name of flag as string>,<default value>,<Short Description of flag>)
	intPtr := flag.Int("int", 10, "test for int")
	boolPtr := flag.Bool("bool", false, "test for bool")

	var svar string
	flag.StringVar(&svar, "svar", "bar", "Predefined String var") // flag.<Datatype>Var(<pre-defined var pointer to store>,<name of flag as string>,<default value>,<Short Description of flag>)

	flag.Parse()

	fmt.Println("word:", *stringPtr)
	fmt.Println("numb:", *intPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())

	/*
		go run cmd_args_w_flags.go -word="test while running" -bool=true -int=69 -svar="summer of 69"
		word: test while running
		numb: 69
		fork: true
		svar: summer of 69
		tail: []
	*/
}

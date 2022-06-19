package main

import (
	"Searcher/searcher"
	"flag"
	"fmt"
	"strings"
)

func main() {
	directory := flag.String("dir", ".", "directory")
	str := flag.String("strings", "", "strings to find")

	flag.Parse()

	input := strings.Split(*str, ",")

	s := searcher.New(input)
	fmt.Printf("%v\n", s.Contain(directory))
}

package main

import (
	"flag"
	"fmt"

	"github.com/playground/httphash/client"
)

func main() {
	var parallel int
	flag.IntVar(&parallel, "parallel", 10, "number of parallel requests")
	flag.Parse()
	urls := flag.Args()

	httphash, err := client.New(parallel, urls)
	if err != nil {
		fmt.Printf("error creating myhttp client; err: %s\n", err)
		return
	}

	httphash.Process()

}

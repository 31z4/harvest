//+build !test

package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr,
			"harvest helps you understand what's inside your Redis.\n\n"+
				"usage: harvest [-samples n] redis_url\n",
		)
		flag.PrintDefaults()
	}

	samples := flag.Uint("samples", 1000, "number of samples")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintf(os.Stderr, "missing required argument: redis_url\n")
		flag.Usage()
		os.Exit(2)
	}

	redisUrl := flag.Arg(0)

	results, err := Sample(redisUrl, int(*samples))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(results)
}

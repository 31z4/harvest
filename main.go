package main

import (
	"fmt"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	kingpin.Version("0.1.0")
	kingpin.CommandLine.Help = "harvest helps you understand what's inside your Redis."

	samples := kingpin.Flag("samples", "maximum number of samples").Short('s').Default("1000").Int()
	results := kingpin.Flag("results", "maximum number of output results").Short('n').Default("10").Int()
	redisUrl := kingpin.Arg("url", "redis URL").Default("redis://localhost").String()

	kingpin.Parse()

	output, err := Sample(*redisUrl, *samples, *results)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(output)
}

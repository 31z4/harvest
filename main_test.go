package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"testing"
)

const binaryName = "build/harvest"

func TestHarvestArguments(t *testing.T) {
	cases := []struct {
		args   []string
		output []byte
	}{
		{
			[]string{},
			[]byte("missing required argument: redis_url"),
		},
		{
			[]string{"-samples"},
			[]byte("flag needs an argument: -samples"),
		},
		{
			[]string{"-samples", "test", "test"},
			[]byte("invalid value \"test\" for flag -samples: strconv.ParseUint: parsing \"test\": invalid syntax"),
		},
		{
			[]string{"-samples", "-1", "test"},
			[]byte("invalid value \"-1\" for flag -samples: strconv.ParseUint: parsing \"-1\": invalid syntax"),
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			command := exec.Command(binaryName, c.args...)
			output, err := command.CombinedOutput()
			firstLine := bytes.Split(output, []byte("\n"))[0]

			if !reflect.DeepEqual(c.output, firstLine) {
				println(err)
				t.Errorf("expected: %#v\nresult: %#v", c.output, output)
			}
		})
	}
}

func TestHarvest(t *testing.T) {
	err := exec.Command(binaryName, "redis://localhost").Run()
	if err != nil {
		t.Errorf("unexpected error %#v", err)
	}
}

func TestMain(m *testing.M) {
	err := exec.Command("make").Run()
	if err != nil {
		fmt.Printf("could not make %s: %v", binaryName, err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}

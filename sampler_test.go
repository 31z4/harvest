package main

import (
	"strconv"
	"testing"
)

func TestSampleArguments(t *testing.T) {
	cases := []struct {
		redisUrl string
		samples  int
		results  int
		err      string
	}{
		{
			"redis://localhost",
			0,
			10,
			"number of samples must be > 0",
		},
		{
			"redis://localhost",
			10,
			0,
			"number of results must be > 0",
		},
		{
			"localhost",
			1,
			1,
			"invalid redis URL scheme: ",
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := Sample(c.redisUrl, c.samples, c.results)
			if err.Error() != c.err {
				t.Errorf("expected: %#v\nresult: %#v", c.err, err)
			}
		})
	}
}

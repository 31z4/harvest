package main

import (
	"strconv"
	"testing"
)

func TestSampleArguments(t *testing.T) {
	cases := []struct {
		redisUrl string
		samples  int
		err      string
	}{
		{
			"redis://localhost",
			0,
			"number of samples must be > 0",
		},
		{
			"localhost",
			1,
			"invalid redis URL scheme: ",
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			_, err := Sample(c.redisUrl, c.samples)
			if err.Error() != c.err {
				t.Errorf("expected: %#v\nresult: %#v", c.err, err)
			}
		})
	}
}

func TestSample(t *testing.T) {
	_, err := Sample("redis://localhost", 1)
	if err != nil {
		t.Errorf("unexpected error %#v", err)
	}
}

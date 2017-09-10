package main

import (
	"strconv"
	"testing"

	"github.com/go-redis/redis"
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

func TestSample(t *testing.T) {
	var client = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// Check is Redis database is empty before running tests
	dbSize, err := client.DBSize().Result()
	if err != nil {
		t.Errorf("DBSIZE: error: %#v\n", err)
	}

	if dbSize > 0 {
		t.Error("redis database is not empty")
	}

	t.Run("empty database", func(t *testing.T) {
		output, err := Sample("redis://localhost:6379", 1, 1)
		if err == nil {
			t.Errorf("expected an error but got: %#v", output)
		}
	})

	t.Run("non empty database", func(t *testing.T) {
		err := client.Set("key", "value", 0).Err()
		if err != nil {
			t.Errorf("SET: error: %#v\n", err)
		}

		output, err := Sample("redis://localhost:6379", 10, 1)
		if err != nil {
			t.Errorf("unexpected error: %#v", err)
		}

		if output != "key: 100.00% (10)" {
			t.Errorf("unexpected output: %#v", output)
		}
	})

	// Flush Redis database after running tests
	err = client.FlushDB().Err()
	if err != nil {
		t.Errorf("FLUSHDB: error: %#v\n", err)
	}
}

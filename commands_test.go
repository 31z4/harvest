package main

import (
	"testing"

	"github.com/go-redis/redis"
)

func TestMemoryUsage(t *testing.T) {
	var client = redis.NewClient(&redis.Options{Addr: "localhost:6379"})

	// Check is Redis database is empty before running tests
	dbSize, err := client.DBSize().Result()
	if err != nil {
		t.Errorf("DBSIZE: error: %#v\n", err)
	}

	if dbSize > 0 {
		t.Error("redis database is not empty")
	}

	t.Run("nonexistent key", func(t *testing.T) {
		v, err := MemoryUsage(client, "nonexistent").Result()
		if err == nil {
			t.Errorf("expected an error but got: %#v", v)
		}
	})

	t.Run("existing key", func(t *testing.T) {
		const value = "value"

		err := client.Set("key", value, 0).Err()
		if err != nil {
			t.Errorf("SET: error: %#v\n", err)
		}

		v, err := MemoryUsage(client, "key").Result()
		if err != nil {
			t.Errorf("MEMORY USAGE: error: %#v\n", err)
		}

		const expected = int64(56)
		if v != expected {
			t.Errorf("expected %#v, got %#v", expected, v)
		}
	})

	// Flush Redis database after running tests
	err = client.FlushDB().Err()
	if err != nil {
		t.Errorf("FLUSHDB: error: %#v\n", err)
	}
}

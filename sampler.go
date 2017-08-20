package main

import (
	"errors"

	"github.com/go-redis/redis"
)

// Sample samples Redis keys and returns statistics about it.
func Sample(redisUrl string, samples uint) error {
	if samples == 0 {
		return errors.New("number of samples must be > 0")
	}

	_, err := redis.ParseURL(redisUrl)
	if err != nil {
		return err
	}

	return nil
}

package main

import (
	"errors"

	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// Sample samples Redis keys and returns statistics about it.
func Sample(redisUrl string, samples, results int) (string, error) {
	if samples <= 0 {
		return "", errors.New("number of samples must be > 0")
	}

	if results <= 0 {
		return "", errors.New("number of results must be > 0")
	}

	opt, err := redis.ParseURL(redisUrl)
	if err != nil {
		return "", err
	}

	client := redis.NewClient(opt)

	var dbSize int64
	dbSize, err = client.DBSize().Result()
	if err != nil {
		return "", err
	}

	if dbSize == 0 {
		return "", errors.New("the database is empty")
	}

	if dbSize < int64(samples) {
		fmt.Fprintf(os.Stderr, "warning: database size (%v) is less than the number of samples (%v)\n\n", dbSize, samples)
	}

	trie := NewTrie()

	for i := 0; i < samples; i++ {
		key, err := client.RandomKey().Result()
		if err != nil {
			return "", err
		}

		trie.Insert(key)
	}

	return trie.Sprint(results), nil
}

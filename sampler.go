package main

import (
	"errors"

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

	trie := NewTrie()
	client := redis.NewClient(opt)

	for i := 0; i < samples; i++ {
		key, err := client.RandomKey().Result()
		if err != nil {
			return "", err
		}

		trie.Insert(key)
	}

	return trie.Sprint(results), nil
}

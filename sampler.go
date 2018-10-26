package main

import (
	"errors"

	"fmt"
	"os"

	"github.com/go-redis/redis"
)

// Max number of keys to be considered as small
const thresholdDbSize = 1000000

// COUNT argument to redis's SCAN command
const scanCount = 1000

type Strategy interface {
	handle(*Trie, *redis.Client, int) error
}

// RandomStrategy implements the sampling of keys using redis's RANDOMKEY command.
type RandomStrategy struct{}

func (RandomStrategy) handle(trie *Trie, client *redis.Client, samples int) error {
	for i := 0; i < samples; i++ {
		key, err := client.RandomKey().Result()
		if err != nil {
			return err
		}

		trie.Insert(key)
	}
	return nil
}

// SmallDbStrategy implements the sampling using redis's SCAN command.
type SmallDbStrategy struct{}

func (SmallDbStrategy) handle(trie *Trie, client *redis.Client, samples int) error {
	// we will ignore the argument "samples"
	var cursor uint64
	for {
		var keys []string
		var err error
		keys, cursor, err = client.Scan(cursor, "", scanCount).Result()
		if err != nil {
			return err
		}
		for _, key := range keys {
			trie.Insert(key)
		}
		if cursor == 0 {
			break
		}
	}
	return nil
}

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

	var strategy Strategy
	if dbSize <= int64(thresholdDbSize) {
		strategy = SmallDbStrategy{}
	} else {
		strategy = RandomStrategy{}
	}
	err = strategy.handle(trie, client, samples)
	if err != nil {
		return "", err
	}

	return trie.Sprint(results), nil
}

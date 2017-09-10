package main

import "github.com/go-redis/redis"

// MemoryUsage implements Redis MEMORY USAGE command.
func MemoryUsage(cleint *redis.Client, key string) *redis.IntCmd {
	cmd := redis.NewIntCmd("memory", "usage", key)
	cleint.Process(cmd)
	return cmd
}

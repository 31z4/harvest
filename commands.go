package main

import "github.com/go-redis/redis"

// MemoryUsage implements Redis MEMORY USAGE command.
func MemoryUsage(client *redis.Client, key string) *redis.IntCmd {
	cmd := redis.NewIntCmd("memory", "usage", key)
	client.Process(cmd)
	return cmd
}

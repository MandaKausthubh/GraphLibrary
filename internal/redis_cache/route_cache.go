package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/redis/go-redis/v9"
	"github.com/MandaKausthubh/GraphLibrary/internal/graph"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

// NewRedisClient initializes a new Redis connection
func NewRedisClient(addr string, password string, db int) *RedisClient {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password, // "" if no password
		DB:       db,
	})
	return &RedisClient{Client: rdb}
}

// --- PATH CACHING ---

func (r *RedisClient) CachePath(fromID, toID string, data interface{}, ttl time.Duration) error {
	key := fmt.Sprintf("path:%s:%s", fromID, toID)
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, bytes, ttl).Err()
}

func (r *RedisClient) GetCachedPath(fromID, toID string, result interface{}) (bool, error) {
	key := fmt.Sprintf("path:%s:%s", fromID, toID)
	val, err := r.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return false, nil // cache miss
	} else if err != nil {
		return false, err
	}

	return true, json.Unmarshal([]byte(val), result)
}

// --- EDGE CACHING ---

func (r *RedisClient) CacheEdge(edge *graph.Edge, ttl time.Duration) error {
	key := fmt.Sprintf("edge:%s:%s", edge.FromNodeID, edge.ToNodeID)
	bytes, err := json.Marshal(edge)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, bytes, ttl).Err()
}

func (r *RedisClient) GetCachedEdge(fromID, toID string) (*graph.Edge, error) {
	key := fmt.Sprintf("edge:%s:%s", fromID, toID)
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var edge graph.Edge
	if err := json.Unmarshal([]byte(val), &edge); err != nil {
		return nil, err
	}
	return &edge, nil
}

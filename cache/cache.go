// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/predicate"
	"github.com/hazelcast/hazelcast-go-client/serialization"
)

type Cache[T any] struct {
	ctx    context.Context
	client *hazelcast.Client
}

func NewCache[T any](config hazelcast.Config) (*Cache[T], error) {
	var ctx = context.Background()

	client, err := hazelcast.StartNewClientWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &Cache[T]{ctx: ctx, client: client}, nil
}

func NewCacheWithClient[T any](client *hazelcast.Client) *Cache[T] {
	var ctx = context.Background()
	return &Cache[T]{ctx, client}
}

func (c *Cache[T]) Put(mapName string, key string, value T) error {
	mp, err := c.client.GetMap(c.ctx, mapName)
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return mp.Set(c.ctx, key, serialization.JSON(bytes))
}

func (c *Cache[T]) Get(mapName string, key string) (*T, error) {
	mp, err := c.client.GetMap(c.ctx, mapName)
	if err != nil {
		return nil, err
	}

	value, err := mp.Get(c.ctx, key)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	hzJsonValue, ok := value.(serialization.JSON)
	if !ok {
		return nil, fmt.Errorf("value of cached object with key '%s' is not a HazelcastJsonValue", key)
	}

	var unmarshalledValue = new(T)
	if err := json.Unmarshal(hzJsonValue, unmarshalledValue); err != nil {
		return nil, err
	}

	return unmarshalledValue, nil
}

func (c *Cache[T]) Delete(mapName string, key string) error {
	mp, err := c.client.GetMap(c.ctx, mapName)
	if err != nil {
		return err
	}

	err = mp.Delete(c.ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (c *Cache[T]) GetQuery(mapName string, query predicate.Predicate) ([]T, error) {
	mp, err := c.client.GetMap(c.ctx, mapName)
	if err != nil {
		return nil, err
	}

	entries, err := mp.GetEntrySetWithPredicate(c.ctx, query)
	if err != nil {
		return nil, err
	}

	var unmarshalledValues = make([]T, 0)
	for _, entry := range entries {
		hzJsonValue, ok := entry.Value.(serialization.JSON)
		if !ok {
			return nil, fmt.Errorf("value of cached object with key '%s' is not a HazelcastJsonValue", entry.Key.(string))
		}

		var unmarshalledValue T
		if err := json.Unmarshal(hzJsonValue, &unmarshalledValue); err != nil {
			return nil, err
		}

		unmarshalledValues = append(unmarshalledValues, unmarshalledValue)
	}

	return unmarshalledValues, nil
}

func (c *Cache[T]) GetClient() *hazelcast.Client {
	return c.client
}

func (c *Cache[T]) GetMap(mapName string) (*hazelcast.Map, error) {
	return c.client.GetMap(c.ctx, mapName)
}

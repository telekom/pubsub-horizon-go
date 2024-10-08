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

type HazelcastCache[T any] struct {
	ctx    context.Context
	client *hazelcast.Client
}

type HazelcastBasedCache[T any] interface {
	Cache[T]
	GetQuery(mapName string, query predicate.Predicate) ([]T, error)
	GetClient() *hazelcast.Client
	GetMap(mapKey string) (*hazelcast.Map, error)
	AddListener(mapName string, listener Listener[T]) error
}

func NewHazelcastCache[T any](config hazelcast.Config) (*HazelcastCache[T], error) {
	var ctx = context.Background()

	client, err := hazelcast.StartNewClientWithConfig(ctx, config)
	if err != nil {
		return nil, err
	}

	return &HazelcastCache[T]{ctx: ctx, client: client}, nil
}

func NewHazelcastCacheWithClient[T any](client *hazelcast.Client) *HazelcastCache[T] {
	var ctx = context.Background()
	return &HazelcastCache[T]{ctx, client}
}

func (c *HazelcastCache[T]) Put(mapName string, key string, value T) error {
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

func (c *HazelcastCache[T]) Get(mapName string, key string) (*T, error) {
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

	var unmarshalledValue = new(T)
	if err := c.unmarshalHazelcastJson(key, value, unmarshalledValue); err != nil {
		return nil, err
	}

	return unmarshalledValue, nil
}

func (c *HazelcastCache[T]) Delete(mapName string, key string) error {
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

func (c *HazelcastCache[T]) GetQuery(mapName string, query predicate.Predicate) ([]T, error) {
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

func (c *HazelcastCache[T]) GetClient() *hazelcast.Client {
	return c.client
}

func (c *HazelcastCache[T]) GetMap(mapName string) (*hazelcast.Map, error) {
	return c.client.GetMap(c.ctx, mapName)
}

func (c *HazelcastCache[T]) AddListener(mapName string, listener Listener[T]) error {
	mp, err := c.client.GetMap(c.ctx, mapName)
	if err != nil {
		return err
	}

	// Add a listener to the map to react to events.
	_, err = mp.AddListener(c.ctx, hazelcast.MapListener{
		EntryAdded: func(event *hazelcast.EntryNotified) {
			var obj T
			if err := c.unmarshalHazelcastJson(event.Key, event.Value, &obj); err != nil {
				listener.OnError(event, err)
				return
			}

			listener.OnAdd(event, obj)
		},
		EntryUpdated: func(event *hazelcast.EntryNotified) {
			var obj T
			if err := c.unmarshalHazelcastJson(event.Key, event.Value, &obj); err != nil {
				listener.OnError(event, err)
				return
			}

			var oldObj T
			if err := c.unmarshalHazelcastJson(event.Key, event.OldValue, &oldObj); err != nil {
				listener.OnError(event, err)
				return
			}

			listener.OnUpdate(event, obj, oldObj)
		},
		EntryRemoved: func(event *hazelcast.EntryNotified) {
			listener.OnDelete(event)
		},
	}, true)

	if err != nil {
		return fmt.Errorf("failed to add listener: %w", err)
	}

	return nil
}

func (c *HazelcastCache[T]) unmarshalHazelcastJson(key any, value any, obj any) error {
	hzJsonValue, ok := value.(serialization.JSON)
	if !ok {
		return fmt.Errorf("value of cached object with key '%s' is not a HazelcastJsonValue", key)
	}
	return json.Unmarshal(hzJsonValue, obj)
}

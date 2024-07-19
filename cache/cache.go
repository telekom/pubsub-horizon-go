// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cache

import "context"

type Cache[T any] interface {
	Put(mapName string, key string, value T) error
	PutWithContext(ctx context.Context, mapName string, key string, value T) error
	Get(mapName string, key string) (*T, error)
	GetWithContext(ctx context.Context, mapName string, key string) (*T, error)
	Delete(mapName string, key string) error
	DeleteWithContext(ctx context.Context, mapName string, key string) error
	AddListener(mapName string, listener Listener[T]) error
}

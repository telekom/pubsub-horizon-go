// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/cluster"
	"github.com/hazelcast/hazelcast-go-client/predicate"
	"github.com/stretchr/testify/assert"
	"github.com/telekom/pubsub-horizon-go/test"
	"os"
	"testing"
	"time"
)

var cache *HazelcastCache[TestDummy]

type TestDummy struct {
	Foo string `json:"foo"`
}

func TestMain(m *testing.M) {
	test.StartDocker()
	code := m.Run()
	test.StopDocker()
	os.Exit(code)
}

func TestNewCache(t *testing.T) {
	var assertions = assert.New(t)

	var config = hazelcast.Config{}
	config.Cluster.Name = "horizon"
	config.Cluster.Network.SetAddresses(test.GetHazelcastHost())
	config.Cluster.ConnectionStrategy.ReconnectMode = cluster.ReconnectModeOff
	config.Failover.TryCount = 5

	var err error
	cache, err = NewHazelcastCache[TestDummy](config)
	assertions.Nil(err)
}

func TestNewCacheWithClient(t *testing.T) {
	var assertions = assert.New(t)
	var cacheWithSameClient = NewHazelcastCacheWithClient[TestDummy](cache.client)
	assertions.Equal(cache.client, cacheWithSameClient.client)
}

func TestCache_Put(t *testing.T) {
	var assertions = assert.New(t)
	var dummy = TestDummy{
		Foo: "bar",
	}

	assertions.NoError(cache.Put("testMap", "dummy", dummy))
}

func TestCache_Get(t *testing.T) {
	var assertions = assert.New(t)
	dummy, err := cache.Get("testMap", "dummy")
	assertions.NoError(err)
	assertions.Equal("bar", dummy.Foo)
}

func TestCache_GetQuery(t *testing.T) {
	var assertions = assert.New(t)
	var query = predicate.Equal("foo", "bar")

	results, err := cache.GetQuery("testMap", query)
	assertions.NoError(err)
	assertions.Equal(1, len(results))
	assertions.Equal("bar", results[0].Foo)
}

func TestCache_Delete(t *testing.T) {
	var assertions = assert.New(t)

	var dummy = TestDummy{
		Foo: "bar",
	}
	err := cache.Put("testMap", "dummy", dummy)
	assertions.NoError(err)

	err = cache.Delete("testMap", "dummy")
	assertions.NoError(err)

	deletedDummy, err := cache.Get("testMap", "dummy")
	assertions.Nil(deletedDummy)
}

func TestCache_AddListener(t *testing.T) {
	var assertions = assert.New(t)

	var listener = &MockListener[TestDummy]{}
	var listenerDummy = TestDummy{Foo: "bar"}

	var err = cache.AddListener("testMap", listener)
	assertions.NoError(err)

	t.Run("Add", func(t *testing.T) {
		var assertions = assert.New(t)
		var err = cache.Put("testMap", "listenerDummy", listenerDummy)
		assertions.NoError(err)
		assertions.Eventually(func() bool {
			return listener.onAddCalled
		}, 5*time.Second, 10*time.Millisecond)
	})

	t.Run("Update", func(t *testing.T) {
		var assertions = assert.New(t)
		listenerDummy.Foo = "fizz"

		var err = cache.Put("testMap", "listenerDummy", listenerDummy)
		assertions.NoError(err)

		assertions.Eventually(func() bool {
			return listener.onUpdateCalled
		}, 5*time.Second, 10*time.Millisecond)
	})

	t.Run("Delete", func(t *testing.T) {
		var assertions = assert.New(t)

		var err = cache.Delete("testMap", "listenerDummy")
		assertions.NoError(err)
		assertions.Eventually(func() bool {
			return listener.onDeleteCalled
		}, 5*time.Second, 10*time.Millisecond)
	})

	t.Run("Errors", func(t *testing.T) {
		assertions.False(listener.onErrorCalled)
		assertions.NoError(listener.err)
	})
}

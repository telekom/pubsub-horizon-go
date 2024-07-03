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
)

var cache *Cache[TestDummy]

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
	cache, err = NewCache[TestDummy](config)
	assertions.Nil(err)
}

func TestNewCacheWithClient(t *testing.T) {
	var assertions = assert.New(t)
	var cacheWithSameClient = NewCacheWithClient[TestDummy](cache.client)
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

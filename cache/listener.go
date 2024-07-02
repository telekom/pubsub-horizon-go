package cache

import "github.com/hazelcast/hazelcast-go-client"

type Listener[T any] interface {
	OnAdd(event *hazelcast.EntryNotified, obj T)
	OnUpdate(event *hazelcast.EntryNotified, obj T)
	OnDelete(event *hazelcast.EntryNotified, obj T)
}

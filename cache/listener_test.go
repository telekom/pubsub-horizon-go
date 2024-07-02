package cache

import (
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/telekom/pubsub-horizon-go/resource"
)

type MockListener[T resource.SubscriptionResource] struct {
	onAddCalled    bool
	onUpdateCalled bool
	onDeleteCalled bool
}

func (m *MockListener[T]) OnAdd(event *hazelcast.EntryNotified, obj TestDummy) {
	m.onAddCalled = true
}

func (m *MockListener[T]) OnUpdate(event *hazelcast.EntryNotified, obj TestDummy) {
	m.onUpdateCalled = true
}

func (m *MockListener[T]) OnDelete(event *hazelcast.EntryNotified, obj TestDummy) {
	m.onDeleteCalled = true
}

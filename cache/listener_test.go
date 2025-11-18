// Copyright 2025 Deutsche Telekom AG
//
// SPDX-License-Identifier: Apache-2.0

package cache

import (
	"github.com/hazelcast/hazelcast-go-client"
)

type MockListener[T TestDummy] struct {
	onAddCalled    bool
	onUpdateCalled bool
	onDeleteCalled bool
	onErrorCalled  bool
	err            error
}

// event, obj
func (m *MockListener[T]) OnAdd(_ *hazelcast.EntryNotified, _ TestDummy) {
	m.onAddCalled = true
}

// event, obj, oldObj
func (m *MockListener[T]) OnUpdate(_ *hazelcast.EntryNotified, _ TestDummy, _ TestDummy) {
	m.onUpdateCalled = true
}

// event
func (m *MockListener[T]) OnDelete(_ *hazelcast.EntryNotified) {
	m.onDeleteCalled = true
}

// event
func (m *MockListener[T]) OnError(_ *hazelcast.EntryNotified, err error) {
	m.onErrorCalled = true
	m.err = err
}

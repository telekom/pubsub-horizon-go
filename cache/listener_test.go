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

func (m *MockListener[T]) OnAdd(event *hazelcast.EntryNotified, obj TestDummy) {
	_, _ = event, obj
	m.onAddCalled = true
}

func (m *MockListener[T]) OnUpdate(event *hazelcast.EntryNotified, obj TestDummy, oldObj TestDummy) {
	_, _, _ = event, obj, oldObj
	m.onUpdateCalled = true
}

func (m *MockListener[T]) OnDelete(event *hazelcast.EntryNotified) {
	_ = event
	m.onDeleteCalled = true
}

func (m *MockListener[T]) OnError(event *hazelcast.EntryNotified, err error) {
	_ = event
	m.onErrorCalled = true
	m.err = err
}

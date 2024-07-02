// Copyright 2024 Deutsche Telekom IT GmbH
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
	m.onAddCalled = true
}

func (m *MockListener[T]) OnUpdate(event *hazelcast.EntryNotified, obj TestDummy, oldObj TestDummy) {
	m.onUpdateCalled = true
}

func (m *MockListener[T]) OnDelete(event *hazelcast.EntryNotified) {
	m.onDeleteCalled = true
}

func (m *MockListener[T]) OnError(event *hazelcast.EntryNotified, err error) {
	m.onErrorCalled = true
	m.err = err
}

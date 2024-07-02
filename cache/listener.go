// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package cache

import "github.com/hazelcast/hazelcast-go-client"

type Listener[T any] interface {
	OnAdd(event *hazelcast.EntryNotified, obj T)
	OnUpdate(event *hazelcast.EntryNotified, obj T, oldObj T)
	OnDelete(event *hazelcast.EntryNotified)
	OnError(event *hazelcast.EntryNotified, err error)
}

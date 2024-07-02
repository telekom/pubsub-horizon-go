// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

//go:build testing

package test

import (
	"context"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/cluster"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"time"
)

var (
	pool               *dockertest.Pool
	hazelcastContainer *dockertest.Resource

	hazelcastImage = EnvOrDefault("HAZELCAST_IMAGE", "hazelcast/hazelcast")
	hazelcastTag   = EnvOrDefault("HAZELCAST_TAG", "5.3.6")
	hazelcastHost  = EnvOrDefault("HAZELCAST_HOST", "localhost")
	hazelcastPort  = EnvOrDefault("HAZELCAST_PORT", "5701")
)

func StartDocker() {
	var err error

	pool, err = dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	if err = pool.Client.Ping(); err != nil {
		log.Fatalf("Could not ping docker: %s", err)
	}

	pool.MaxWait = 30 * time.Second

	setupHazelcast()

	err = pool.Retry(func() error {
		return pingHazelcast()
	})

	if err != nil {
		log.Fatalf("Could not reach hazelcast after several tries: %s", err)
	}
}

func StopDocker() {
	var err = pool.Purge(hazelcastContainer)
	if err != nil {
		log.Fatalf("Could not purge hazelcast container: %s", err)
	}
}

func setupHazelcast() {
	var err error
	hazelcastContainer, err = pool.RunWithOptions(&dockertest.RunOptions{
		Name:         "horizon-go-hazelcast",
		Repository:   hazelcastImage,
		Tag:          hazelcastTag,
		ExposedPorts: []string{"5701/tcp"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5701/tcp": {{HostIP: hazelcastHost, HostPort: hazelcastPort}},
		},
		Env: []string{
			"HZ_CLUSTERNAME=horizon",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true
		config.RestartPolicy = docker.RestartPolicy{
			Name: "no",
		}
	})

	if err != nil {
		log.Fatalf("Could not create hazelcast container:  %s", err)
	}
}

func pingHazelcast() error {
	var ctx = context.Background()
	config := hazelcast.NewConfig()

	config.Cluster.Name = "horizon"
	config.Cluster.Network.SetAddresses(hazelcastHost)
	config.Cluster.ConnectionStrategy.ReconnectMode = cluster.ReconnectModeOff

	config.Failover.TryCount = 5

	client, err := hazelcast.StartNewClientWithConfig(ctx, config)
	if err != nil {
		log.Printf("Could not connect to hazelcast: %s\n", err)
		return err
	}

	return client.Shutdown(ctx)
}

func GetHazelcastHost() string {
	return hazelcastHost
}

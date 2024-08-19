// Copyright 2024 Deutsche Telekom IT GmbH
//
// SPDX-License-Identifier: Apache-2.0

package util

import (
	"github.com/hazelcast/hazelcast-go-client/logger"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type HazelcastZerologLogger struct {
	level zerolog.Level
}

func NewHazelcastZerologLogger(level zerolog.Level) *HazelcastZerologLogger {
	return &HazelcastZerologLogger{level}
}

func (l *HazelcastZerologLogger) Log(weight logger.Weight, f func() string) {
	var messageLevel = l.translateWeight(weight)
	if messageLevel >= l.level {
		log.WithLevel(messageLevel).Msgf("Hazelcast: %s", f())
	}
}

func (*HazelcastZerologLogger) translateWeight(weight logger.Weight) zerolog.Level {
	switch weight {

	case logger.WeightDebug, logger.WeightTrace:
		return zerolog.DebugLevel

	case logger.WeightInfo:
		return zerolog.InfoLevel

	case logger.WeightWarn:
		return zerolog.WarnLevel

	case logger.WeightError:
		return zerolog.ErrorLevel

	case logger.WeightFatal:
		return zerolog.FatalLevel

	default:
		return zerolog.InfoLevel
	}
}
